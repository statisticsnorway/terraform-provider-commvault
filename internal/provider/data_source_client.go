package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"net/http"
	"statisticsnorway/terraform-provider-commvault/pkg/commvault/apiclient"
	"strconv"
)

type clientDataSource struct {
	api *apiclient.APIClient
}

type clientDataSourceModel struct {
	Id   types.String `tfsdk:"client_id"`
	Name types.String `tfsdk:"client_name"`
}

var (
	_ datasource.DataSource              = &clientDataSource{}
	_ datasource.DataSourceWithConfigure = &clientDataSource{}
)

func NewClientDataSource() datasource.DataSource {
	return &clientDataSource{}
}

func (d *clientDataSource) Configure(ctx context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	providerData, ok := request.ProviderData.(*CommvaultProviderData)
	if !ok {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Please report this issue to the provider developers.",
		)

		return
	}

	d.api = providerData.ApiClient
}

func (d *clientDataSource) Metadata(ctx context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {

	response.TypeName = request.ProviderTypeName + "_client"
}

func (d *clientDataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_id": schema.StringAttribute{
				Optional: true,
			},
			"client_name": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (d *clientDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data clientDataSourceModel

	// Read Terraform configuration data into the model
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	id := data.Id.ValueString()
	name := data.Name.ValueString()

	if id == "" && name == "" {
		response.Diagnostics.AddError("One of `client_id` or `client_name` must be set", "")
		return
	}
	if id != "" && name != "" {
		response.Diagnostics.AddError("Only one of `client_id` or `client_name` should be set", "")
		return
	}

	var (
		clientResponse apiclient.ClientGetByIdResponse
		_              *http.Response
		err            error
	)

	if id != "" {
		clientResponse, _, err = d.api.ClientApi.GetById(ctx, id)
		if err != nil {
			response.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get client by ID, got error: %s", err))
			return
		}
	} else {
		clientResponse, _, err = d.api.ClientApi.GetByName(ctx, name)
		if err != nil {
			response.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get client by name, got error: %s", err))
			return
		}
	}

	if len(clientResponse.ClientProperties) == 0 {
		response.Diagnostics.AddError("Could not found client", "Unable to get client. Empty response")
		return
	}

	clientEntity := clientResponse.ClientProperties[0].Client.ClientEntity
	data.Id = types.StringValue(strconv.Itoa(clientEntity.ClientId))
	data.Name = types.StringValue(clientEntity.ClientName)

	diags := response.State.Set(ctx, &data)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}
