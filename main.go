package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/cyrilgdn/terraform-provider-postgresql/postgresql"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: postgresql.Provider})
}
