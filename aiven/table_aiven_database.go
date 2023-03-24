package aiven

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAivenDatabase(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_database",
		Description: "Retrieve information about your databases.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"project_name", "service_name"}),
			Hydrate:    listDatabases,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project_name", "service_name", "database_name"}),
			Hydrate:    getDatabase,
		},
		Columns: []*plugin.Column{
			{
				Name:        "database_name",
				Type:        proto.ColumnType_STRING,
				Description: "The database name.",
			},
			{
				Name:        "project_name",
				Type:        proto.ColumnType_STRING,
				Description: "The project name.",
			},
			{
				Name:        "service_name",
				Type:        proto.ColumnType_STRING,
				Description: "The service name.",
			},
			{
				Name:        "lc_collate",
				Type:        proto.ColumnType_STRING,
				Description: "Default string sort order (LC_COLLATE) for PostgreSQL database.",
			},
			{
				Name:        "lc_type",
				Type:        proto.ColumnType_STRING,
				Description: "Default character classification (LC_CTYPE) for PostgreSQL database.",
			},
		},
	}
}

type AivenDatabase struct {
	ProjectName string
	ServiceName string
	aiven.Database
}

func listDatabases(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	project := d.EqualsQuals["project_name"].GetStringValue()
	service := d.EqualsQuals["service_name"].GetStringValue()

	// Check if project or service is empty.
	if project == "" || service == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listDatabases", "connection_error", err)
		return nil, err
	}

	databaseList, err := conn.Databases.List(project, service)
	if err != nil {
		plugin.Logger(ctx).Error("listDatabases", "api_error", err)
		return nil, err
	}

	for _, database := range databaseList {
		d.StreamListItem(ctx, AivenDatabase{project, service, *database})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getDatabase(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	project := d.EqualsQuals["project_name"].GetStringValue()
	service := d.EqualsQuals["service_name"].GetStringValue()
	database := d.EqualsQuals["database_name"].GetStringValue()

	// Check if project or service or database is empty.
	if project == "" || service == "" || database == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getDatabase", "connection_error", err)
		return nil, err
	}

	data, err := conn.Databases.Get(project, service, database)
	if err != nil {
		plugin.Logger(ctx).Error("getDatabase", "api_error", err)
		return nil, err
	}

	return AivenDatabase{project, service, *data}, nil
}
