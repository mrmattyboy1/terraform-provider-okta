package okta

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-okta/sdk"
	"github.com/okta/terraform-provider-okta/sdk/query"
)

type GroupAssignmentModel struct {
	Id       string `json:"id" mapstructure:"id"`
	Priority int64  `json:"priority" mapstructure:"priority"`
	Profile  string `json:"profile" mapstructure:"profile"`
}

func dataSourceAppGroupAssignments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppGroupAssignmentsRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the Okta App being queried for groups",
				ForceNew:    true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "A group to associate with the application",
						},
						"priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Priority of group assignment",
						},
						"profile": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)",
						},
					},
				},
				Description: "List of groups IDs assigned to the app",
			},
		},
		Description: "Get a set of groups assigned to an Okta application.",
	}
}

// type GroupAssignmentData struct {
// 	id       string
// 	priority int64
// 	profile  string
// }

func dataSourceAppGroupAssignmentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := getOktaClientFromMetadata(m)
	id := d.Get("id").(string)

	groupAssignments, resp, err := client.Application.ListApplicationGroupAssignments(ctx, id, &query.Params{})
	if err != nil {
		return diag.Errorf("unable to query for groups from app (%s): %s", id, err)
	}

	for {
		var moreAssignments []*sdk.ApplicationGroupAssignment
		if resp.HasNextPage() {
			resp, err = resp.Next(ctx, &moreAssignments)
			if err != nil {
				return diag.Errorf("unable to query for groups from app (%s): %s", id, err)
			}
			groupAssignments = append(groupAssignments, moreAssignments...)
		} else {
			break
		}
	}

	var groups []GroupAssignmentModel
	for _, assignment := range groupAssignments {

		profileBytes, err := json.Marshal(assignment.Profile)
		if err != nil {
			continue
		}

		groupAssignment := GroupAssignmentModel{
			Id:       assignment.Id,
			Profile:  string(profileBytes[:]),
			Priority: assignment.Priority,
		}

		groups = append(groups, groupAssignment)
		fmt.Printf("Matt: %v", assignment.Profile)
	}
	_ = d.Set("groups", groups)
	d.SetId(id)
	return nil
}
