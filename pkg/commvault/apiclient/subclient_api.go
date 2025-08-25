package apiclient

import (
	"context"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type SubclientApiService service

// Get
// This operation returns a list of subclients for a client.
// https://api.commvault.com/docs/SP36/api/cv/SubclientOperations/get-subclient-id/
func (a *SubclientApiService) Get(ctx context.Context, subClientId string) (SubclientGetResponse, *http.Response, error) {
	return prepareAndCallApiJSON[SubclientGetResponse](ctx, a.client, http.MethodGet, "/subclient", url.Values{
		"client_id": {subClientId},
	}, nil)
}

// Create a new subclient
// https://api.commvault.com/docs/SP36/api/cv/SubclientOperations/post-subclient/
func (a *SubclientApiService) Create(ctx context.Context) (SubclientGetResponse, *http.Response, error) {
	return prepareAndCallApiJSON[SubclientGetResponse](ctx, a.client, http.MethodPost, "/subclient", nil, nil)
}

// Update Subclient
// https://api.commvault.com/docs/SP38/api/cv/SubclientOperations/update-subclient-properties/#update-subclient-properties
func (a *SubclientApiService) Update(ctx context.Context, subclientId string, createRequest *SubclientUpdateRequest) (interface{}, *http.Response, error) {
	return prepareAndCallApiJSON[interface{}](ctx, a.client, http.MethodPost, "/subclient/"+subclientId, nil, createRequest)
}
