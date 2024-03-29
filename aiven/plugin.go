package aiven

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Plugin creates this (aiven) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-aiven",
		DefaultTransform: transform.FromCamel(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"aiven_account":                      tableAivenAccount(ctx),
			"aiven_account_authentication":       tableAivenAccountAuthentication(ctx),
			"aiven_account_team":                 tableAivenAccountTeam(ctx),
			"aiven_billing_group":                tableAivenBillingGroup(ctx),
			"aiven_project":                      tableAivenProject(ctx),
			"aiven_project_event_log":            tableAivenProjectEventLog(ctx),
			"aiven_project_user":                 tableAivenProjectUser(ctx),
			"aiven_project_vpc":                  tableAivenProjectVpc(ctx),
			"aiven_service":                      tableAivenService(ctx),
			"aiven_service_integration_endpoint": tableAivenServiceIntegrationEndpoint(ctx),
		},
	}
	return p
}
