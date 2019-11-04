package slack

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema {
			"token": {
				Type: schema.TypeString,
				Required: true,
				Description: "Slack API access token",

			},
		},
		DataSourcesMap: map[string]*schema.Resource {
			"slack_get_users": dataSourceSlackGetUsers(),
		},
		ResourcesMap: map[string]*schema.Resource {

		},
		ConfigureFunc: providerConfigure,
	}

	return p
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var client *client_api

	if token, ok := d.GetOk("token"); ok {
		client.token = token.(string)
	} else {
		return nil, fmt.Errorf("missing token field for the Slack provider")
	}

	return &client, nil
}