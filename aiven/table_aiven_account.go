package aiven

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAivenAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_account",
		Description: "Retrieve information about your accounts.",
		List: &plugin.ListConfig{
			Hydrate: listAccounts,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAccount,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the account.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the account.",
			},
			{
				Name:        "billing_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the billing is enabled for the account.",
			},
			{
				Name:        "owner_team_id",
				Type:        proto.ColumnType_STRING,
				Description: "The owner team ID of the account.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The create time of the account.",
			},
			{
				Name:        "update_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The update time of the account.",
			},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: "The tenant ID of the account.",
			},
		},
	}
}

func listAccounts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAccounts", "connection_error", err)
		return nil, err
	}

	accountList, err := conn.Accounts.List()
	if err != nil {
		plugin.Logger(ctx).Error("listAccounts", "api_error", err)
		return nil, err
	}

	for _, account := range accountList.Accounts {
		d.StreamListItem(ctx, account)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getAccount", "connection_error", err)
		return nil, err
	}

	id := d.EqualsQuals["id"].GetStringValue()

	// Check if id is empty.
	if id == "" {
		return nil, nil
	}

	accountList, err := conn.Accounts.Get(id)
	if err != nil {
		plugin.Logger(ctx).Error("getAccount", "api_error", err)
		return nil, err
	}

	return accountList.Account, nil
}
