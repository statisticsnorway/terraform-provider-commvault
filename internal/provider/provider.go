package provider

import (
	"context"
	"os"

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

type providerModel struct {
	BaseURL  types.String `tfsdk:"base_url"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *commvaultProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "commvaultx"
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
	var cfg providerModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	api := &APIClient{
		BaseURL:  cfg.BaseURL.ValueString(),
		Username: first(cfg.Username.ValueString(), os.Getenv("COMMVAULT_USERNAME")),
		Password: first(cfg.Password.ValueString(), os.Getenv("COMMVAULT_PASSWORD")),
		HTTP:     NewHTTPClientInsecure(),
	}

	token, err := api.Login(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Login failed", err.Error())
		return
	}
	api.Token = token
	resp.ResourceData = api
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
