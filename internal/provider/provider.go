package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"os"
	"statisticsnorway/terraform-provider-commvault/pkg/commvault/apiclient"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = (*commvaultProvider)(nil)

func New(version string) func() provider.Provider {
	return func() provider.Provider { return &commvaultProvider{version: version} }
}

type commvaultProvider struct{ version string }

type commvaultProviderModel struct {
	BaseURL  types.String `tfsdk:"base_url"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type CommvaultProviderData struct {
	ApiClient *apiclient.APIClient
}

func (p *commvaultProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "commvault"
	resp.Version = p.version
}

func (p *commvaultProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{Required: true},
			"username": schema.StringAttribute{Optional: true},
			"password": schema.StringAttribute{Optional: true, Sensitive: true},
		},
	}
}

func (p *commvaultProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config commvaultProviderModel
	tflog.Info(ctx, "Configuring client")
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate config
	if config.BaseURL.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Unknown Commvault API Base URL",
			"The provider cannot create the Commvault API client as there is an unknown configuration value for the Commvault API base url. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the COMMVAULT_BASE_URL environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Commvault API Username",
			"The provider cannot create the Commvault API client as there is an unknown configuration value for the Commvault API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the COMMVAULT_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Commvault API Password",
			"The provider cannot create the Commvault API client as there is an unknown configuration value for the Commvault API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the COMMVAULT_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	basePath := first(config.BaseURL.ValueString(), os.Getenv("COMMVAULT_BASE_URL"))
	apiClient := apiclient.NewAPIClient(&apiclient.Configuration{
		BasePath:   basePath,
		UserAgent:  "terraform-provider-commvault@" + p.version,
		HTTPClient: NewHTTPClient(120*time.Second, true),
	})

	tflog.Info(ctx, fmt.Sprintf("Login to API at: %s", basePath))
	loginResponse, _, err := apiClient.LoginApi.Login(ctx, &apiclient.LoginRequest{
		Username: first(config.Username.ValueString(), os.Getenv("COMMVAULT_USERNAME")),
		Password: first(config.Password.ValueString(), os.Getenv("COMMVAULT_PASSWORD")),
	})
	if err != nil {
		resp.Diagnostics.AddError("Login failed", err.Error())
		return
	}
	tflog.Info(ctx, "Login OK")
	apiClient.SetToken(loginResponse.Token)

	providerData := &CommvaultProviderData{
		ApiClient: apiClient,
	}
	resp.ResourceData = providerData
	resp.DataSourceData = providerData
}

func (p *commvaultProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{NewClientResource}
}

func (p *commvaultProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func first(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
