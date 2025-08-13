package commvault

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Base API URL, e.g. https://<commvault-host>/commandcenter/api",
				DefaultFunc: schema.EnvDefaultFunc("COMMVAULT_BASE_URL", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Commvault username",
				DefaultFunc: schema.EnvDefaultFunc("COMMVAULT_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Commvault password",
				DefaultFunc: schema.EnvDefaultFunc("COMMVAULT_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"commvault_client": resourceCommvaultClient(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return &CommvaultConfig{
				BaseURL:  d.Get("base_url").(string),
				Username: d.Get("username").(string),
				Password: d.Get("password").(string),
			}, nil
		},
	}
}
