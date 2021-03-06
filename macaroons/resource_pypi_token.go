package macaroons

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"gopkg.in/macaroon.v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ProjectCaveatPermissions struct {
	Projects []string `json:"projects"`
}

type ProjectCaveat struct {
	Version     int                      `json:"version"`
	Permissions ProjectCaveatPermissions `json:"permissions"`
}

func resourcePypiToken() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a method of restricting user-scoped pypi.org API tokens into projects-scoped tokens.",
		CreateContext: resourcePypiTokenCreate,
		UpdateContext: resourcePypiTokenCreate,
		ReadContext:   resourcePypiTokenNoOp,
		DeleteContext: resourcePypiTokenNoOp,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "A user-scoped API token from pypi.org.",
			},
			"project": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of project to create a project-scoped token for.",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The resulting API token that can only be used for given project.",
			},
		},
	}
}

func resourcePypiTokenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sourceTokenBytes, err := macaroon.Base64Decode([]byte(d.Get("source_token").(string)[5:]))
	if err != nil {
		return diag.FromErr(err)
	}
	var token macaroon.Macaroon
	err = token.UnmarshalBinary(sourceTokenBytes)
	if err != nil {
		return diag.FromErr(err)
	}
	caveat, err := json.Marshal(ProjectCaveat{
		Version: 1,
		Permissions: ProjectCaveatPermissions{
			Projects: []string{d.Get("project").(string)},
		},
	})
	if err != nil {
		return diag.FromErr(err)
	}
	token.AddFirstPartyCaveat(caveat)
	tokenBytes, err := token.MarshalBinary()
	if err != nil {
		return diag.FromErr(err)
	}
	id := sha1.Sum(tokenBytes)
	d.SetId(base64.StdEncoding.EncodeToString(id[:]))
	d.Set("token", "pypi-"+base64.StdEncoding.EncodeToString([]byte(tokenBytes)))
	return nil
}

func resourcePypiTokenNoOp(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
