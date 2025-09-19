package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
	"statisticsnorway/terraform-provider-commvault/pkg/commvault/apiclient"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = (*clientResource)(nil)
	_ resource.ResourceWithImportState = (*clientResource)(nil)
)

func NewClientResource() resource.Resource { return &clientResource{} }

type clientResource struct {
	api *apiclient.APIClient
}

type bucketItemModel struct {
	Name    types.String `tfsdk:"name"`
	Project types.String `tfsdk:"project"`
}

type clientModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	PlanID       types.Int64  `tfsdk:"plan_id"`
	CredentialID types.Int64  `tfsdk:"credential_id"`
	AccessNodeID types.Int64  `tfsdk:"access_node_id"`
	ProjectID    types.String `tfsdk:"project_id"`

	BucketContents []bucketItemModel `tfsdk:"bucket_contents"`

	SubclientID types.Int64 `tfsdk:"subclient_id"`
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
			"credential_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"access_node_id": schema.Int64Attribute{
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
				Computed: true,
			},
		},
		Blocks: map[string]schema.Block{
			"bucket_contents": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"project": schema.StringAttribute{
							Required: true,
						},
					},
				},
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
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Please report this issue to the provider developers.")
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

	tflog.Info(ctx, "Creating client")
	createResponse, response, err := r.api.ClientApi.Create(ctx, buildCreateClientPayload(plan))
	if err != nil {
		if response != nil && response.StatusCode == 409 {
			resp.Diagnostics.AddError("Client already exists", fmt.Sprintf("Http status code: %d", response.StatusCode))
			return
		}
		resp.Diagnostics.AddError("Create failed", err.Error())
		return
	}

	clientId := strconv.Itoa(createResponse.Response.Entity.ClientID)
	tflog.Debug(ctx, "New client created id:"+clientId)

	if len(plan.BucketContents) == 0 {
		resp.State.Set(ctx, &clientModel{
			ID:             types.StringValue(clientId),
			Name:           plan.Name,
			PlanID:         plan.PlanID,
			CredentialID:   plan.CredentialID,
			AccessNodeID:   plan.AccessNodeID,
			ProjectID:      plan.ProjectID,
			SubclientID:    plan.SubclientID,
			BucketContents: plan.BucketContents,
		})
		return
	}

	var subID string
	if !plan.SubclientID.IsNull() && plan.SubclientID.ValueInt64() > 0 {
		subID = strconv.Itoa(int(plan.SubclientID.ValueInt64()))
		tflog.Debug(ctx, "Using subclient id from plan: "+subID)
	} else {
		tflog.Debug(ctx, "Fetching default subclient id for clientId "+clientId)
		getResponse, httpResponse, err2 := r.api.SubclientApi.Get(ctx, clientId)
		if err2 != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed fetching subclient for client id: %s", clientId), err2.Error())
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("get subclient response: %#v", getResponse))
		if httpResponse.StatusCode == http.StatusNotFound || len(getResponse.SubClientProperties) == 0 || getResponse.SubClientProperties[0].SubClientEntity.SubclientId == 0 {
			resp.Diagnostics.AddError("Failed fetching subclient", fmt.Sprintf("Client id %s. HTTP %d", clientId, httpResponse.StatusCode))
			return
		}
		subID = strconv.Itoa(getResponse.SubClientProperties[0].SubClientEntity.SubclientId)
	}
	tflog.Info(ctx, "Using subclient "+subID)

	payload := buildCloudAppsSubclientPayload(subID, plan.BucketContents)

	_, httpResponse, err := r.api.SubclientApi.Update(ctx, subID, payload)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to update subclientId: %d", subID), err.Error())
		return
	}
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Failed fetching subclient", fmt.Sprintf("Client id: %s, subclientId: %d, Http status: %d", clientId, subID, httpResponse.StatusCode))
		return
	}
	tflog.Info(ctx, "Subclient updated")

	subclientId, _ := strconv.Atoi(subID)
	plan.SubclientID = types.Int64Value(int64(subclientId))
	diags := resp.State.Set(ctx, &clientModel{
		ID:             types.StringValue(clientId),
		Name:           plan.Name,
		PlanID:         plan.PlanID,
		CredentialID:   plan.CredentialID,
		AccessNodeID:   plan.AccessNodeID,
		ProjectID:      plan.ProjectID,
		SubclientID:    plan.SubclientID,
		BucketContents: plan.BucketContents,
	})
	resp.Diagnostics.Append(diags...)
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

func (r *clientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan clientModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	var state clientModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	subclientId := strconv.Itoa(int(state.SubclientID.ValueInt64()))
	tflog.Debug(ctx, "Updating subclient id "+subclientId)
	payload := buildCloudAppsSubclientPayload(subclientId, plan.BucketContents)
	tflog.Debug(ctx, "New payload built. Sending update request")
	_, httpResponse, err := r.api.SubclientApi.Update(ctx, subclientId, payload)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to update subclientId: %s", subclientId), err.Error())
		return
	}
	// Check if response is OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to update subclientId: %s", subclientId), fmt.Sprintf("Status code  %s: %s", httpResponse.StatusCode))
		return
	}

	// Check if bucket contents has changed

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state clientModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || state.ID.IsNull() {
		return
	}

	clientId := state.ID.ValueString()
	_, httpResponse, err := r.api.ClientApi.Delete(ctx, clientId, true)
	if err != nil {
		resp.Diagnostics.AddError("Error occurred during delete", err.Error())
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

func buildCloudAppsSubclientPayload(subclientID string, contents []bucketItemModel) *apiclient.SubclientCreateOrUpdateRequestAndResponse {
	bucketPaths := make([]apiclient.SubclientFsContent, len(contents))
	for _, c := range contents {
		bucketPaths = append(bucketPaths, apiclient.SubclientFsContent{
			Path: "/" + c.Name.ValueString(),
		})
	}

	gcpContents := make([]apiclient.Content, len(contents))
	for _, item := range contents {
		gcpContents = append(gcpContents, apiclient.Content{
			GCPContent: &apiclient.GCPContent{
				BucketName:  item.Name.ValueString(),
				ProjectName: item.Project.ValueString(),
			},
		})
	}

	subclientId, _ := strconv.Atoi(subclientID)
	return &apiclient.SubclientCreateOrUpdateRequestAndResponse{
		SubclientProperties: apiclient.SubclientProperties{
			SubclientEntity: apiclient.SubclientUpdateRequestClientEntity{
				SubclientID: subclientId,
			},
			UseLocalContent:              true,
			FsContentOperationType:       "OVERWRITE",
			FsExcludeFilterOperationType: "CLEAR",
			FsIncludeFilterOperationType: "CLEAR",
			Content:                      bucketPaths,
			CloudAppsSubClientProp: &apiclient.CloudAppsSubClientProp{
				InstanceType: "GOOGLE_CLOUD",
				ObjectStorageSubclient: apiclient.ObjectStorageSubclient{
					ContentOperationType: "OVERWRITE",
					Content:              gcpContents,
				},
			},
		},
	}
}
