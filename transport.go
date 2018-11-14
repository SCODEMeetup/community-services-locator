package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeProviderEndpoint(svc FoodPantryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.Providers()
		if err != nil {
			return providerResponse{v, err.Error()}, nil
		}
		return providerResponse{v, ""}, nil
	}
}

func decodeProviderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type providerRequest struct {
}

type providerResponse struct {
	Providers []Provider `json:"providers"`
	Err       string     `json:"err,omitempty"`
}
