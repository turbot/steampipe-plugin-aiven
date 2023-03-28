package aiven

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAivenProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_project",
		Description: "Retrieve information about your projects.",
		List: &plugin.ListConfig{
			Hydrate: listProjects,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getProject,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The project name.",
			},
			{
				Name:        "account_id",
				Type:        proto.ColumnType_STRING,
				Description: "The account ID.",
			},

			{
				Name:        "available_credits",
				Type:        proto.ColumnType_STRING,
				Description: "Available credits, in USD.",
			},
			{
				Name:        "billing_extra_text",
				Type:        proto.ColumnType_STRING,
				Description: "Extra text to be included in all project invoices, e.g. purchase order or cost center number.",
			},
			{
				Name:        "country",
				Type:        proto.ColumnType_STRING,
				Description: "The country.",
			},
			{
				Name:        "country_code",
				Type:        proto.ColumnType_STRING,
				Description: "Two letter ISO country code.",
			},
			{
				Name:        "default_cloud",
				Type:        proto.ColumnType_STRING,
				Description: "Default cloud to use when launching services.",
			},
			{
				Name:        "estimated_balance",
				Type:        proto.ColumnType_STRING,
				Description: "Estimated balance, in USD.",
			},
			{
				Name:        "payment_method",
				Type:        proto.ColumnType_STRING,
				Description: "The payment method.",
			},
			{
				Name:        "vat_id",
				Type:        proto.ColumnType_STRING,
				Description: "EU VAT Identification Number.",
				Transform:   transform.FromField("VatID"),
			},
			{
				Name:        "billing_currency",
				Type:        proto.ColumnType_STRING,
				Description: "The billing currency.",
			},
			{
				Name:        "copy_from_project",
				Type:        proto.ColumnType_STRING,
				Description: "Project name from which to copy settings to the new project.",
			},
			{
				Name:        "billing_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "Billing group ID.",
			},
			{
				Name:        "billing_group_name",
				Type:        proto.ColumnType_STRING,
				Description: "Billing group name.",
			},
			{
				Name:        "billing_emails",
				Type:        proto.ColumnType_JSON,
				Description: "List of project billing email addresses.",
			},
			{
				Name:        "card",
				Type:        proto.ColumnType_JSON,
				Description: "Credit card assigned to the project.",
			},
			{
				Name:        "technical_emails",
				Type:        proto.ColumnType_JSON,
				Description: "List of technical email addresses.",
			},
		},
	}
}

func listProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listProjects", "connection_error", err)
		return nil, err
	}

	projectList, err := conn.Projects.List()
	if err != nil {
		plugin.Logger(ctx).Error("listProjects", "api_error", err)
		return nil, err
	}

	for _, project := range projectList {
		d.StreamListItem(ctx, project)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	name := d.EqualsQuals["name"].GetStringValue()

	// Check if name is empty
	if name == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getProject", "connection_error", err)
		return nil, err
	}

	project, err := conn.Projects.Get(name)
	if err != nil {
		plugin.Logger(ctx).Error("getProject", "api_error", err)
		return nil, err
	}

	return project, nil
}
