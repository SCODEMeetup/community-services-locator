package main

import (
	"time"
)

// FoodPantryService provides operations on strings.
type FoodPantryService interface {
	Providers() ([]Provider, error)
}

type Provider struct {
	ID                 int
	Code               string
	Description        string
	Level              int
	Active             string
	DHSFlag            string
	Text               string
	BypassFollowupFlag string
	VolunteerFlag      string
	AddUser            string
	AddDate            time.Time
	SubCatID           int
	DHSDescription     string
	UpdateUser         string
	UpdateDate         time.Time
}

type foodPantryService struct {
}

func (f foodPantryService) Providers() ([]Provider, error) {
	return []Provider{
		{
			ID:          1,
			Description: "Test Provider",
		},
	}, nil
}
