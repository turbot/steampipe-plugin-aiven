package aiven

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAivenAccountTeamMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_account_team_member",
		Description: "Retrieve information about your account team members.",
		List: &plugin.ListConfig{
			Hydrate: listAccountTeamMembers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "account_id",
					Require:    plugin.Required,
					CacheMatch: "exact",
				},
				{
					Name:       "team_id",
					Require:    plugin.Required,
					CacheMatch: "exact",
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "user_id",
				Type:        proto.ColumnType_STRING,
				Description: "The user ID.",
			},
			{
				Name:        "real_name",
				Type:        proto.ColumnType_STRING,
				Description: "User real name.",
			},
			{
				Name:        "user_email",
				Type:        proto.ColumnType_STRING,
				Description: "User email address.",
			},
			{
				Name:        "account_id",
				Type:        proto.ColumnType_STRING,
				Description: "The account ID.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The create time of the team member.",
			},
			{
				Name:        "team_id",
				Type:        proto.ColumnType_STRING,
				Description: "The team ID.",
			},
			{
				Name:        "team_name",
				Type:        proto.ColumnType_STRING,
				Description: "The team name.",
			},
			{
				Name:        "update_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The update time of the team member.",
			},
		},
	}
}

type AccountTeamMember struct {
	AccountId string
	aiven.AccountTeamMember
}

func listAccountTeamMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	accountId := d.EqualsQuals["account_id"].GetStringValue()
	teamId := d.EqualsQuals["team_id"].GetStringValue()

	// Check if accountId or teamId is empty.
	if accountId == "" || teamId == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAccountTeamMembers", "connection_error", err)
		return nil, err
	}

	memberList, err := conn.AccountTeamMembers.List(accountId, teamId)
	if err != nil {
		plugin.Logger(ctx).Error("listAccountTeamMembers", "api_error", err)
		return nil, err
	}

	for _, member := range memberList.Members {
		d.StreamListItem(ctx, AccountTeamMember{accountId, member})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
