package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	_ resource.ResourceWithImportState = (*clientResource)(nil)
)

func NewClientResource() resource.Resource { return &clientResource{} }

type clientResource struct{ api *APIClient }

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
			"bucket_name": schema.StringAttribute{
				Optional: true,
			},
			"bucket_project": schema.StringAttribute{
				Optional: true,
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

func (r *clientResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData != nil {
		r.api = req.ProviderData.(*APIClient)
	}
}

func (r *clientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan clientModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.api.doJSON(ctx, http.MethodPost, r.api.BaseURL+"/Client", buildPayload(plan))
	if err != nil {
		resp.Diagnostics.AddError("Create failed", err.Error())
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		b, _ := io.ReadAll(httpResp.Body)
		resp.Diagnostics.AddError("Create failed", fmt.Sprintf("HTTP %d: %s", httpResp.StatusCode, b))
		return
	}

	var cr createResponse
	_ = json.NewDecoder(httpResp.Body).Decode(&cr)
	id := strconv.Itoa(cr.Response.Entity.ClientID)

	if !plan.BucketName.IsNull() && plan.BucketName.ValueString() != "" {
		bucket := plan.BucketName.ValueString()
		proj := plan.BucketProject.ValueString()
		if proj == "" {
			proj = plan.ProjectID.ValueString()
		}

		var subID int64
		if !plan.SubclientID.IsNull() && plan.SubclientID.ValueInt64() > 0 {
			subID = plan.SubclientID.ValueInt64()
		} else {
			var errFind error
			subID, errFind = r.api.findDefaultSubclientID(ctx, id)
			if errFind != nil {
				resp.Diagnostics.AddWarning("Bucket binding: discovery error", errFind.Error())
			}
			if subID == 0 {
				resp.Diagnostics.AddWarning(
					"Bucket binding skipped",
					"Could not discover default subclient. Provide 'subclient_id' to bind the bucket.")
			}
		}
		if subID > 0 {
			if err := r.api.bindGCSBucket(ctx, subID, bucket, proj); err != nil {
				resp.Diagnostics.AddError("Bucket binding failed", err.Error())
			} else {
				resp.Diagnostics.AddWarning(
					"Bucket bound",
					fmt.Sprintf("Bucket %s was successfully bound to subclient %d", bucket, subID),
				)
			}
		}
	}

	resp.State.Set(ctx, &clientModel{
		ID:            types.StringValue(id),
		Name:          plan.Name,
		PlanID:        plan.PlanID,
		CredentialID:  plan.CredentialID,
		AccessNodeID:  plan.AccessNodeID,
		ProjectID:     plan.ProjectID,
		BucketName:    plan.BucketName,
		BucketProject: plan.BucketProject,
		SubclientID:   plan.SubclientID,
		Response:      types.StringValue(fmt.Sprintf(`{"clientId":%s,"clientName":"%s"}`, id, plan.Name.ValueString())),
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

	httpResp, err := r.api.doJSON(ctx, http.MethodGet, fmt.Sprintf("%s/Client/%s", r.api.BaseURL, state.ID.ValueString()), nil)
	if err != nil || httpResp.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	defer httpResp.Body.Close()
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
	_, _ = r.api.doJSON(ctx, http.MethodDelete, fmt.Sprintf("%s/Client/%s?forceDelete=true", r.api.BaseURL, state.ID.ValueString()), nil)
}

func (r *clientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func buildPayload(plan clientModel) map[string]any {
	return map[string]any{
		"clientInfo": map[string]any{
			"clientType": 15,
			"cloudClonnectorProperties": map[string]any{
				"instance": map[string]any{
					"cloudAppsInstance": map[string]any{
						"credentialType": "GOOGLE_SERVICE_ACCOUNT",
						"generalCloudProperties": map[string]any{
							"credentials": map[string]any{
								"credentialId": plan.CredentialID.ValueInt64(),
							},
							"memberServers": []map[string]any{{
								"client": map[string]any{
									"_type_":   3,
									"clientId": plan.AccessNodeID.ValueInt64(),
								},
							}},
							"numberOfBackupStreams": 1,
						},
						"googleCloudInstance": map[string]any{
							"GCPProjectId": plan.ProjectID.ValueString(),
							"serverName":   "storage.googleapis.com",
						},
						"instanceType":            "GOOGLE_CLOUD",
						"instanceTypeDisplayName": "Google Cloud",
					},
					"instance": map[string]any{
						"applicationId": 134,
						"commCellId":    2,
						"instanceName":  plan.Name.ValueString(),
					},
				},
				"instanceType": "GOOGLE_CLOUD",
			},
			"plan": map[string]any{
				"planId": plan.PlanID.ValueInt64(),
			},
		},
		"entity": map[string]any{
			"clientName": plan.Name.ValueString(),
		},
	}
}
