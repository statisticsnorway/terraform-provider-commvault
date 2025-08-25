package apiclient

import (
	"context"
	"net/http"
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
// https://api.commvault.com/docs/SP36/api/cv/ClientOperations/get-client/
func (a *ClientApiService) GetById(ctx context.Context, clientId string) (ClientGetResponse, *http.Response, error) {
	return prepareAndCallApiJSON[ClientGetResponse](ctx, a.client, http.MethodGet, "/client/"+clientId, nil, nil)
}

// Create a new client
// TODO: Find documentation for endpoint
func (a *ClientApiService) Create(ctx context.Context, createRequest *ClientCreateRequest) (ClientCreateResponse, *http.Response, error) {
	//body := &ClientCreateRequest{
	//	ClientInfo: ClientInfo{
	//		ClientType: 15,
	//		CloudConnectorProperties: CloudConnectorProperties{
	//			Instance: Instance{
	//				CloudAppsInstance: CloudAppsInstance{
	//					CredentialType: "GOOGLE_SERVICE_ACCOUNT",
	//					GeneralCloudProperties: GeneralCloudProperties{
	//						Credentials: Credentials{
	//							CredentialID: plan.CredentialId.ValueInt64(),
	//						},
	//						MemberServers: []MemberServer{
	//							MemberServer{
	//								Client: ClientCreateRequestClient{
	//									Type:     3,
	//									ClientID: plan.AccessNodeID.ValueInt64(),
	//								},
	//							},
	//						},
	//						NumberOfBackupStreams: 1,
	//					},
	//					GoogleCloudInstance: GoogleCloudInstance{
	//						GCPProjectID: plan.ProjectID.ValueString(),
	//						ServerName:   "storage.googleapis.com",
	//					},
	//					InstanceType:            "GOOGLE_CLOUD",
	//					InstanceTypeDisplayName: "Google Cloud",
	//				},
	//				Instance: InstanceDetails{
	//					ApplicationID: 123,
	//					CommCellID:    2,
	//					InstanceName:  plan.Name.ValueString(),
	//				},
	//			},
	//			InstanceType: "GOOGLE_CLOUD",
	//		},
	//		Plan: ClientCreateRequestPlan{
	//			PlanID: plan.PlanID.ValueInt64(),
	//		},
	//	},
	//	Entity: Entity{
	//		ClientName: plan.Name.ValueString(),
	//	},
	//}

	return prepareAndCallApiJSON[ClientCreateResponse](ctx, a.client, http.MethodPost, "/client", nil, createRequest)
}

// Update
// Updates a client https://api.commvault.com/docs/SP38/api/cv/ClientOperations/post-client-client-id/
