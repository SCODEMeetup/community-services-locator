package main

import (
	"encoding/json"
	"io/ioutil"
	http "net/http"
)

// TaxonomyService provides operations on strings.
type TaxonomyService interface {
	Taxonomy() ([]Record, error)
}

type Record struct {
	ID                 int         `json:"_id"`
	TAXONID            int         `json:"TAXON_ID"`
	TAXONOMYCODE       string      `json:"TAXONOMY_CODE"`
	DESCRIPTION        string      `json:"DESCRIPTION"`
	TAXONOMYLEVEL      int         `json:"TAXONOMY_LEVEL"`
	ACTIVEFLAG         string      `json:"ACTIVE_FLAG"`
	DHSFLAG            string      `json:"DHS_FLAG"`
	TEXT               string      `json:"TEXT"`
	BYPASSFOLLOWUPFLAG string      `json:"BYPASS_FOLLOWUP_FLAG"`
	VOLUNTEERFLAG      string      `json:"VOLUNTEER_FLAG"`
	ADDUSER            string      `json:"ADD_USER"`
	ADDDATE            string      `json:"ADD_DATE"`
	TAXONIDSUBCATOF    interface{} `json:"TAXON_ID_SUBCAT_OF"`
	DHSDESCRIPTION     string      `json:"DHS_DESCRIPTION"`
	UPDATEUSER         string      `json:"UPDATE_USER"`
	UPDATEDATE         string      `json:"UPDATE_DATE"`
}

type Taxonomy struct {
	Help    string `json:"help"`
	Success bool   `json:"success"`
	Result  struct {
		IncludeTotal bool   `json:"include_total"`
		ResourceID   string `json:"resource_id"`
		Fields       []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"fields"`
		RecordsFormat string   `json:"records_format"`
		Records       []Record `json:"records"`
		Limit         int      `json:"limit"`
		Links         struct {
			Start string `json:"start"`
			Next  string `json:"next"`
		} `json:"_links"`
		Total int `json:"total"`
	} `json:"result"`
}

type taxonomyService struct {
}

func (f taxonomyService) Taxonomy() ([]Record, error) {
	resp, err := http.Get("https://ckan.smartcolumbusos.com/api/3/action/datastore_search?resource_id=371dd944-411c-4851-a065-9f3f605ddfb9")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	taxonomy := Taxonomy{}
	err = json.Unmarshal(body, &taxonomy)

	if err != nil {
		return nil, err
	}

	return taxonomy.Result.Records, nil
}
