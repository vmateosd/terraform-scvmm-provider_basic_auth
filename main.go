package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/vmateosd/terraform-provider-scvmm/scvmm"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: scvmm.Provider})
}
