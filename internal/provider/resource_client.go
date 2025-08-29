package provider

import (
	"context"
	"fmt"
	"net/http"
	"statisticsnorway/terraform-provider-commvault/pkg/commvault/apiclient"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = (*clientResource)(nil)
	_ resource.ResourceWithConfigure   = (*clientResource)(nil)
	_ resource.ResourceWithImportState = (*clientResource)(nil)
)

func NewClientResource() resource.Resource { return &clientResource{} }

type clientResource struct {
	api *apiclient.APIClient
}

type clientModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	PlanID       types.Int64  `tfsdk:"plan_id"`
	CredentialID types.Int64  `tfsdk:"credential_id"`
	AccessNodeID types.Int64  `tfsdk:"access_node_id"`
	ProjectID    types.String `tfsdk:"project_id"`

	BucketName    types.String `tfsdk:"bucket_name"`
	BucketProject types.String `tfsdk:"bucket_project"`
	SubclientID   types.Int64  `tfsdk:"subclient_id"`
	Response      types.String `tfsdk:"response"`
}

func (r *clientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_client"
}

func (r *clientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"plan_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"project_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"subclient_id": schema.Int64Attribute{
				Optional: true,
			},

			"response": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *clientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*CommvaultProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Please report this issue to the provider developers.",
		)

		return
	}

	r.api = providerData.ApiClient
}

func (r *clientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan clientModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createResponse, _, err := r.api.ClientApi.Create(ctx, buildCreateClientPayload(plan))
	if err != nil {
		resp.Diagnostics.AddError("Create failed", err.Error())
		return
	}

	clientId := strconv.Itoa(createResponse.Response.Entity.ClientID)

	resp.State.Set(ctx, &clientModel{
		ID:           types.StringValue(clientId),
		Name:         plan.Name,
		PlanID:       plan.PlanID,
		CredentialID: plan.CredentialID,
		AccessNodeID: plan.AccessNodeID,
		ProjectID:    plan.ProjectID,
		Response:     types.StringValue(fmt.Sprintf(`{"clientId":%s,"clientName":"%s"}`, clientId, plan.Name.ValueString())),
	})
}

func (r *clientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state clientModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if state.ID.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	clientId := state.ID.ValueString()
	getResponse, httpResp, err := r.api.ClientApi.GetById(ctx, clientId)
	if err != nil || httpResp.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if len(getResponse.ClientProperties) == 0 || strconv.Itoa(getResponse.ClientProperties[0].Client.ClientEntity.ClientId) != clientId {
		gotClientId := strconv.Itoa(getResponse.ClientProperties[0].Client.ClientEntity.ClientId)
		resp.Diagnostics.AddError(fmt.Sprintf("Client id did not match whats in state. Got: %s, expected: %s", gotClientId, clientId), "Client property missing")
	}

	resp.State.Set(ctx, &state)
}

func (r *clientResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "Changes require replacement.")
}

func (r *clientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state clientModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || state.ID.IsNull() {
		return
	}

	//TODO: implement deletion protection, as deletion of client might delete all backup data associated with it as well

	clientId := state.ID.ValueString()
	_, httpResponse, err := r.api.ClientApi.Delete(ctx, clientId, true)
	if err != nil {
		resp.Diagnostics.AddError("Error occured during delete", err.Error())
		return
	}
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Unexpected http status code after delete", httpResponse.Status)
		return
	}
}

func (r *clientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func buildCreateClientPayload(plan clientModel) *apiclient.ClientCreateRequest {
	return &apiclient.ClientCreateRequest{
		ClientInfo: apiclient.ClientInfo{
			ClientType: 15,
			CloudConnectorProperties: apiclient.CloudConnectorProperties{
				Instance: apiclient.Instance{
					CloudAppsInstance: apiclient.CloudAppsInstance{
						CredentialType: "GOOGLE_SERVICE_ACCOUNT",
						GeneralCloudProperties: apiclient.GeneralCloudProperties{
							Credentials: apiclient.Credentials{
								CredentialID: plan.CredentialID.ValueInt64(),
							},
							MemberServers: []apiclient.MemberServer{
								{
									Client: apiclient.ClientCreateRequestClient{
										Type:     3,
										ClientID: plan.AccessNodeID.ValueInt64(),
									},
								},
							},
							NumberOfBackupStreams: 1,
						},
						GoogleCloudInstance: apiclient.GoogleCloudInstance{
							GCPProjectID: plan.ProjectID.ValueString(),
							ServerName:   "storage.googleapis.com",
						},
						InstanceType:            "GOOGLE_CLOUD",
						InstanceTypeDisplayName: "Google Cloud",
					},
					Instance: apiclient.InstanceDetails{
						ApplicationID: 123,
						CommCellID:    2,
						InstanceName:  plan.Name.ValueString(),
					},
				},
				InstanceType: "GOOGLE_CLOUD",
			},
			Plan: apiclient.ClientCreateRequestPlan{
				PlanID: plan.PlanID.ValueInt64(),
			},
		},
		Entity: apiclient.Entity{
			ClientName: plan.Name.ValueString(),
		},
	}
}
