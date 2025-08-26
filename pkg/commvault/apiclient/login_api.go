package apiclient

import (
	"context"
	"net/http"
)

// Linger please
var (
	_ context.Context
)

type LoginApiService service

/*
Login to the commcell using this request and save the token before calling any other API.
https://api.commvault.com/docs/SP36/api/cv/AuthenticationOperations/login/

  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param  *LoginRequest:
  - @param "username" (string) -  Username of the user
  - @param "password" (string) -  A Base64 UTF-8 encoded password for the user

@return LoginResponseCustomerManaged
*/
func (a *LoginApiService) Login(ctx context.Context, requestBody *LoginRequest) (LoginResponseCustomerManaged, *http.Response, error) {
	return prepareAndCallApiJSON[LoginResponseCustomerManaged](ctx, a.client, http.MethodPost, "/login", nil, &requestBody)
}
