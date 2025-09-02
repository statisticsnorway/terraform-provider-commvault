package provider

import (
	"context"
	"encoding/json"
	"fmt"
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

	SubclientID types.Int64  `tfsdk:"subclient_id"`
	Response    types.String `tfsdk:"response"`
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
				Optional: true,
			},
			"response": schema.StringAttribute{
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
							Optional: true,
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

	createResponse, _, err := r.api.ClientApi.Create(ctx, buildCreateClientPayload(plan))
	if err != nil {
		resp.Diagnostics.AddError("Create failed", err.Error())
		return
	}

	clientId := strconv.Itoa(createResponse.Response.Entity.ClientID)

	cloudContents := buildCloudContents(plan)
	if len(cloudContents) > 0 {
		var subID int
		if !plan.SubclientID.IsNull() && plan.SubclientID.ValueInt64() > 0 {
			subID = int(plan.SubclientID.ValueInt64())
		} else {
			getResponse, httpResponse, err := r.api.SubclientApi.Get(ctx, clientId)
			if err != nil {
				resp.Diagnostics.AddError(fmt.Sprintf("Failed fetching subclient for client id: %s", clientId), err.Error())
				return
			}
			if httpResponse.StatusCode == http.StatusNotFound || len(getResponse.SubClientProperties) == 0 {
				resp.Diagnostics.AddError("Failed fetching subclient", fmt.Sprintf("Client id %s. HTTP %d", clientId, httpResponse.StatusCode))
				return
			}
			subID = getResponse.SubClientProperties[0].SubClientEntity.SubclientId
			if subID == 0 {
				resp.Diagnostics.AddError("Failed fetching subclient", "Received subclientId=0")
				return
			}
		}

		payload := buildCloudAppsSubclientPayload(subID, cloudContents)

		if b, err := json.MarshalIndent(payload, "", "  "); err == nil {
			resp.Diagnostics.AddWarning("Subclient Update Payload", string(b))
		}

		_, h, err := r.api.SubclientApi.Update(ctx, strconv.Itoa(subID), payload)
		if err != nil {
			status := ""
			if h != nil {
				status = h.Status
			}
			if ge, ok := err.(apiclient.GenericSwaggerError); ok {
				resp.Diagnostics.AddError(
					fmt.Sprintf("Failed to update subclient %d", subID),
					fmt.Sprintf("HTTP %s\nResponse body:\n%s", status, string(ge.Body())),
				)
				return
			}
			resp.Diagnostics.AddError(
				fmt.Sprintf("Failed to update subclient %d", subID),
				fmt.Sprintf("HTTP %s\n%v", status, err),
			)
			return
		}
	}

	resp.State.Set(ctx, &clientModel{
		ID:             types.StringValue(clientId),
		Name:           plan.Name,
		PlanID:         plan.PlanID,
		CredentialID:   plan.CredentialID,
		AccessNodeID:   plan.AccessNodeID,
		ProjectID:      plan.ProjectID,
		SubclientID:    plan.SubclientID,
		BucketContents: plan.BucketContents,
		Response:       types.StringValue(fmt.Sprintf(`{"clientId":%s,"clientName":"%s"}`, clientId, plan.Name.ValueString())),
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

func buildCloudContents(plan clientModel) []apiclient.Content {
	var contents []apiclient.Content
	for _, item := range plan.BucketContents {
		if item.Name.IsNull() || item.Name.ValueString() == "" {
			continue
		}
		project := plan.ProjectID.ValueString()
		if !item.Project.IsNull() && item.Project.ValueString() != "" {
			project = item.Project.ValueString()
		}
		contents = append(contents, apiclient.Content{
			GCPContent: &apiclient.GCPContent{
				BucketName:  item.Name.ValueString(),
				ProjectName: project,
			},
		})
	}
	return contents
}

func buildCloudAppsSubclientPayload(subclientID int, contents []apiclient.Content) *apiclient.SubclientCreateOrUpdateRequestAndResponse {
	paths := make([]apiclient.SubclientFsContent, 0, len(contents))
	for _, c := range contents {
		if c.GCPContent != nil && c.GCPContent.BucketName != "" {
			paths = append(paths, apiclient.SubclientFsContent{
				Path: "/" + c.GCPContent.BucketName,
			})
		}
	}

	return &apiclient.SubclientCreateOrUpdateRequestAndResponse{
		SubclientProperties: apiclient.SubclientProperties{
			SubclientEntity: apiclient.SubclientUpdateRequestClientEntity{
				SubclientID: subclientID,
			},
			UseLocalContent:              true,
			FsContentOperationType:       "OVERWRITE",
			FsExcludeFilterOperationType: "CLEAR",
			FsIncludeFilterOperationType: "CLEAR",
			Content:                      paths,
			CloudAppsSubClientProp: &apiclient.CloudAppsSubClientProp{
				InstanceType: "GOOGLE_CLOUD",
				ObjectStorageSubclient: apiclient.ObjectStorageSubclient{
					ContentOperationType: "OVERWRITE",
					Content:              contents,
				},
			},
		},
	}
}
