package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/hcevensen/terraform-provider-commvaultx/internal/provider"
)

var version = "dev"

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.Parse()

	if err := providerserver.Serve(
		context.Background(),
		provider.New(version),
		providerserver.ServeOpts{
			// Use your final address once you publish. For local dev,
			// keep this address and configure dev_overrides in ~/.terraformrc
			Address: "hashicorp.com/edu/commvaultx",
			Debug:   debug,
		},
	); err != nil {
		log.Fatal(err)
	}
}
