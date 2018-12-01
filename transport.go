package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeTaxonomyEndpoint(svc TaxonomyService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.Taxonomy()
		if err != nil {
			return taxonomyResponse{v, err.Error()}, nil
		}
		return taxonomyResponse{v, ""}, nil
	}
}

func decodeTaxonomyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type taxonomyRequest struct {
}

type taxonomyResponse struct {
	Records []Record `json:"records"`
	Err     string   `json:"err,omitempty"`
}
