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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = (*subclientResource)(nil)
	_ resource.ResourceWithImportState = (*subclientResource)(nil)
)

func NewSubclientResource() resource.Resource { return &subclientResource{} }

type subclientResource struct {
	api *apiclient.APIClient
}

type subclientModel struct {
	Id          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	ClientName  types.String `tfsdk:"client_id"`
	GcpContents []gcpContent `tfsdk:"gcp_contents"`
}

type gcpContent struct {
	BucketName types.String `tfsdk:"bucket_name"`
	ProjectId  types.String `tfsdk:"project_id"`
}

func (r *subclientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_client"
}

func (r *subclientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"client_id": schema.StringAttribute{
				Required: true,
			},
			"gcp_contents": schema.ListNestedAttribute{
				Required: false,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bucket_name": schema.StringAttribute{
							Required: true,
						},
						"project_id": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}

func (r *subclientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *subclientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan subclientModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	clientName := plan.ClientName.ValueString()
	subclientName := plan.Name.ValueString()
	entity := apiclient.SubclientUpdateRequestClientEntity{
		ClientName:    clientName,
		AppName:       0,
		SubclientName: subclientName,
		BackupsetName: 0,
	}
	content := make([]apiclient.Content, len(plan.GcpContents))
	for _, gcpPlan := range plan.GcpContents {
		content = append(content, apiclient.Content{
			GCPContent: apiclient.GCPContent{
				BucketName:  gcpPlan.BucketName.ValueString(),
				ProjectName: gcpPlan.ProjectId.ValueString(), // Commvault mean projectId not project name..
			},
		})
	}
	_, h, err := r.api.SubclientApi.Create(ctx, buildCreateSubclientPayload(entity, content))
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to create subclient for subclientId: %d, bucket %s, project: %s", subID, bucket, gcpProject), err.Error())
		return
	}
	if h.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to create subclient for subclientId: %d, bucket %s, project: %s, http code: %d", subID, bucket, gcpProject, h.StatusCode), h.Status)
		return
	}

	resp.State.Set(ctx, &subclientModel{
		Id:            plan.Id,
		Name:          plan.Name,
		PlanID:        plan.PlanID,
		BucketName:    plan.BucketName,
		BucketProject: plan.BucketProject,
	})
}

func (r *subclientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state subclientModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if state.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	clientId := state.ClientName.ValueString()
	getResponse, httpResp, err := r.api.ClientApi.GetById(ctx, clientId)
	if err != nil || httpResp.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if len(getResponse.ClientProperties) == 0 || strconv.Itoa(getResponse.ClientProperties[0].Client.ClientEntity.ClientId) != clientId {
		gotClientId := strconv.Itoa(getResponse.ClientProperties[0].Client.ClientEntity.ClientId)
		resp.Diagnostics.AddError(fmt.Sprintf("Client id did not match whats in state. Got: %s, expected: %s", gotClientId, clientId), "Client property missing")
		return
	}

	resp.State.Set(ctx, &state)
}

func (r *subclientResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "Changes require replacement.")
}

func (r *subclientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state subclientModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || state.ID.IsNull() {
		return
	}

	subclientId := state.ID.ValueString()
	_, httpResponse, err := r.api.SubclientApi.Delete(ctx, subclientId)
	if err != nil {
		resp.Diagnostics.AddError("Error occured during delete", err.Error())
		return
	}
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("Unexpected http status code after delete", httpResponse.Status)
		return
	}
}

func (r *subclientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func buildCreateSubclientPayload(subclientEntity apiclient.SubclientUpdateRequestClientEntity, content []apiclient.Content) *apiclient.SubclientCreateOrUpdateRequestAndResponse {
	return &apiclient.SubclientCreateOrUpdateRequestAndResponse{
		SubclientProperties: apiclient.SubclientProperties{
			SubclientEntity: subclientEntity,
			UseLocalContent: true,
			CloudAppsSubClientProp: apiclient.CloudAppsSubClientProp{
				InstanceType: 20,
				ObjectStorageSubclient: apiclient.ObjectStorageSubclient{
					ContentOperationType: 2,
					Content:              content,
				},
			},
		},
	}
}
