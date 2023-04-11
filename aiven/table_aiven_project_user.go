package aiven

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAivenProjectUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_project_user",
		Description: "Retrieve information about your project users.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listProjectUsers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "project_name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project_name", "email"}),
			Hydrate:    getProjectUser,
		},
		Columns: []*plugin.Column{
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "User email address.",
			},
			{
				Name:        "real_name",
				Type:        proto.ColumnType_STRING,
				Description: "User real name.",
			},
			{
				Name:        "billing_contact",
				Type:        proto.ColumnType_BOOL,
				Description: "Set for project's billing contacts.",
			},
			{
				Name:        "member_type",
				Type:        proto.ColumnType_STRING,
				Description: "Project member type.",
			},
			{
				Name:        "project_name",
				Type:        proto.ColumnType_STRING,
				Description: "The project name.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp in ISO 8601 format, always in UTC.",
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
				Name:        "auth_methods",
				Type:        proto.ColumnType_JSON,
				Description: "List of user's required authentication methods.",
			},
		},
	}
}

type ProjectUser struct {
	ProjectName string
	aiven.ProjectUser
}

func listProjectUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*aiven.Project)
	project_name := d.EqualsQuals["project_name"].GetStringValue()

	// check if the provided project_name is not matching with the parentHydrate
	if project_name != "" && project_name != project.Name {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_project_user.listProjectUsers", "connection_error", err)
		return nil, err
	}

	projectUsers, _, err := conn.ProjectUsers.List(project.Name)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_project_user.listProjectUsers", "api_error", err)
		return nil, err
	}

	for _, user := range projectUsers {
		d.StreamListItem(ctx, ProjectUser{project.Name, *user})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getProjectUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	project := d.EqualsQuals["project_name"].GetStringValue()
	email := d.EqualsQuals["email"].GetStringValue()

	// Check if project or email is empty
	if project == "" || email == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_project_user.getProjectUser", "connection_error", err)
		return nil, err
	}

	user, _, err := conn.ProjectUsers.Get(project, email)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_project_user.getProjectUser", "api_error", err)
		return nil, err
	}

	return ProjectUser{project, *user}, nil
}
