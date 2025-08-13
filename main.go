package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hcevensen/terraform-provider-commvaultx/commvault"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: commvault.Provider,
	})
}
