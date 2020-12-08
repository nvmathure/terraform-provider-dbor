package dbor

import (
	"context"
	"database/sql"

	_ "github.com/godror/godror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("DBOR_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("DBOR_PASSWORD", nil),
			},
			"datasource": {
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("DBOR_DATASOURCE", nil),
			},
		},
		ResourcesMap:         map[string]*schema.Resource{},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfig,
	}
}

func providerConfig(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	datasource := d.Get("datasource").(string)

	// Warning or errors can be collected in a slice type
	var diagnostics diag.Diagnostics

	if (username != "") && (password != "") && (datasource != "") {
		c, err := sql.Open("godror", "${username}/${password}@${datasource}")
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diagnostics
	}

	diagnostics = append(diagnostics, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Provider properties username, password and datasource are required",
		Detail:   "Provider properties username, password and datasource are required",
	})

	return nil, diagnostics
}
