package aiven

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAivenBillingGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_billing_group",
		Description: "Retrieve information about your billing groups.",
		List: &plugin.ListConfig{
			Hydrate: listBillingGroups,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getBillingGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the billing group.",
			},
			{
				Name:        "billing_group_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the billing group.",
			},
			{
				Name:        "account_id",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the billing group.",
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "Payment method.",
			},
			{
				Name:        "billing_currency",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the account.",
			},
			{
				Name:        "billing_extra_text",
				Type:        proto.ColumnType_STRING,
				Description: "Available credits, in USD.",
			},
			{
				Name:        "card_id",
				Type:        proto.ColumnType_STRING,
				Description: "Extra text to be included in all billing group invoices, e.g. purchase order or cost center number.",
			},
			{
				Name:        "city",
				Type:        proto.ColumnType_STRING,
				Description: "The country.",
			},
			{
				Name:        "company",
				Type:        proto.ColumnType_STRING,
				Description: "Default cloud to use when launching billing groups.",
			},
			{
				Name:        "country_code",
				Type:        proto.ColumnType_STRING,
				Description: "Estimated balance, in USD.",
			},
			{
				Name:        "vat_id",
				Type:        proto.ColumnType_STRING,
				Description: "EU VAT Identification Number.",
			},
			{
				Name:        "zip_code",
				Type:        proto.ColumnType_STRING,
				Description: "Billing currency.",
			},
			{
				Name:        "billing_emails",
				Type:        proto.ColumnType_JSON,
				Description: "List of project billing email addresses.",
			},
		},
	}
}

func listBillingGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listBillingGroups", "connection_error", err)
		return nil, err
	}

	groupList, err := conn.BillingGroup.ListAll()
	if err != nil {
		plugin.Logger(ctx).Error("listBillingGroups", "api_error", err)
		return nil, err
	}

	for _, group := range groupList {
		d.StreamListItem(ctx, group)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getBillingGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQuals["id"].GetStringValue()

	// Check if id is empty.
	if id == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getBillingGroup", "connection_error", err)
		return nil, err
	}

	group, err := conn.BillingGroup.Get(id)
	if err != nil {
		plugin.Logger(ctx).Error("getBillingGroup", "api_error", err)
		return nil, err
	}

	return group, nil
}
