package apiclient

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// Linger please
var (
	_ context.Context
)

type ClientApiService service

// Get
// This operation returns a list of clients.
// https://api.commvault.com/docs/SP36/api/cv/ClientOperations/get-client/
func (a *ClientApiService) Get(ctx context.Context) (ClientGetResponse, *http.Response, error) {
	return prepareAndCallApiJSON[ClientGetResponse](ctx, a.client, http.MethodGet, "/client/", nil, nil)
}

// GetById
// This operation returns a given client.
// https://api.commvault.com/docs/SP36/api/cv/ClientOperations/get-client-client-id/
func (a *ClientApiService) GetById(ctx context.Context, clientId string) (ClientGetByIdResponse, *http.Response, error) {
	return prepareAndCallApiJSON[ClientGetByIdResponse](ctx, a.client, http.MethodGet, "/client/"+clientId, nil, nil)
}

// Create a new client
// TODO: Find documentation for endpoint
func (a *ClientApiService) Create(ctx context.Context, createRequest *ClientCreateRequest) (ClientCreateResponse, *http.Response, error) {
	return prepareAndCallApiJSON[ClientCreateResponse](ctx, a.client, http.MethodPost, "/client", nil, createRequest)
}

// Delete
// https://api.commvault.com/docs/SP38/api/cv/ClientOperations/delete-client-client-id/
// This operation deletes a client entity.
// This operation also deletes any backup data that is associated with the client.
func (a *ClientApiService) Delete(ctx context.Context, clientId string, forceDelete bool) (ClientDeleteResponse, *http.Response, error) {
	return prepareAndCallApiJSON[ClientDeleteResponse](ctx, a.client, http.MethodDelete, "/client/"+clientId, url.Values{
		"force_delete": {strconv.FormatBool(forceDelete)},
	}, nil)
}
