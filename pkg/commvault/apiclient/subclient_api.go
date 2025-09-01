package apiclient

import (
	"context"
	"net/http"
	"net/url"
)

var (
	_ context.Context
)

type SubclientApiService service

func (a *SubclientApiService) Get(ctx context.Context, subClientId string) (SubclientGetResponse, *http.Response, error) {
	return prepareAndCallApiJSON[SubclientGetResponse](ctx, a.client, http.MethodGet, "/subclient", url.Values{
		"clientId": {subClientId},
	}, nil)
}

func (a *SubclientApiService) Create(ctx context.Context, createRequest *SubclientCreateOrUpdateRequestAndResponse) (SubclientCreateOrUpdateRequestAndResponse, *http.Response, error) {
	return prepareAndCallApiJSON[SubclientCreateOrUpdateRequestAndResponse](ctx, a.client, http.MethodPost, "/subclient", nil, createRequest)
}

func (a *SubclientApiService) Update(ctx context.Context, subclientId string, updateRequest *SubclientCreateOrUpdateRequestAndResponse) (SubclientCreateOrUpdateRequestAndResponse, *http.Response, error) {
	return prepareAndCallApiJSON[SubclientCreateOrUpdateRequestAndResponse](ctx, a.client, http.MethodPost, "/subclient/"+subclientId, nil, updateRequest)
}
