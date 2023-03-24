package aiven

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAivenAccountTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_account_team",
		Description: "Retrieve information about your account teams.",
		List: &plugin.ListConfig{
			ParentHydrate: listAccounts,
			Hydrate:       listAccountTeams,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"account_id", "id"}),
			Hydrate:    getAccountTeam,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the account team.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the account team.",
			},
			{
				Name:        "account_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the account team.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The create time of the account team.",
			},
			{
				Name:        "update_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The update time of the account team.",
			},
		},
	}
}

func listAccountTeams(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	account := h.Item.(aiven.Account)

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAccountTeams", "connection_error", err)
		return nil, err
	}

	teamList, err := conn.AccountTeams.List(account.Id)
	if err != nil {
		plugin.Logger(ctx).Error("listAccountTeams", "api_error", err)
		return nil, err
	}

	for _, team := range teamList.Teams {
		d.StreamListItem(ctx, team)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getAccountTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	account_id := d.EqualsQuals["account_id"].GetStringValue()
	team_id := d.EqualsQuals["id"].GetStringValue()

	// Check if account_id or team_id is empty.
	if account_id == "" || team_id == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getAccountTeam", "connection_error", err)
		return nil, err
	}

	teamList, err := conn.AccountTeams.Get(account_id, team_id)
	if err != nil {
		plugin.Logger(ctx).Error("getAccountTeam", "api_error", err)
		return nil, err
	}

	return teamList.Team, nil
}
