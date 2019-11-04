package main

import "github.com/hashicorp/terraform-plugin-sdk/plugin"
import "github.com/imvalient/terraform-provider-slack"

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: slack.Provider,
	})
}