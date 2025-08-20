package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"statisticsnorway/terraform-provider-commvault/internal/provider"
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
			Address: "hashicorp.com/edu/commvault",
			Debug:   debug,
		},
	); err != nil {
		log.Fatal(err)
	}
}
