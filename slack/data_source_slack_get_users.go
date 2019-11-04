package slack

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSlackGetUsers() *schema.Resource {
	s := &schema.Resource{
		Read: dataSourceSlackGetUsersRead,

		Schema: map[string]*schema.Schema {
			"users": {
				Type: schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema {
						"id": {
							Type: schema.TypeString,
							Computed: true,
						},
						"name": {
							Type: schema.TypeString,
							Computed: true,
						},
						"real_name": {
							Type: schema.TypeString,
							Computed: true,
						},
						"color": {
							Type: schema.TypeString,
							Computed: true,
						},
						"is_bot": {
							Type: schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}

	return s
}

func dataSourceSlackGetUsersRead(d *schema.ResourceData, meta interface{}) error {

	api, err:= meta.(*client_api).Client()
	if err != nil {
		return fmt.Errorf("failed to connect to api slack: %s", err)
	}

	results, err := api.GetUsers()
	if err != nil {
		return err
	}

	users := make([]map[string]interface{}, 0)

	for _, result := range results {

		user := make(map[string]interface{})

		user["id"] = result.ID
		user["name"] = result.Name
		user["real_name"] = result.RealName
		user["color"] = result.Color
		user["is_bot"] = result.IsBot

		users = append(users, user)
	}

	if err := d.Set("users", users); err != nil {
		return fmt.Errorf("Error fetching all users: %+v", err )
	}

	return nil
}

