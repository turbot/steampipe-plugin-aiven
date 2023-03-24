package aiven

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAivenService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_service",
		Description: "Retrieve information about your services.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listServices,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "project_name"}),
			Hydrate:    getService,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The service name.",
			},
			{
				Name:        "project_name",
				Type:        proto.ColumnType_STRING,
				Description: "The project name.",
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "The state of the service.",
			},
			{
				Name:        "plan",
				Type:        proto.ColumnType_STRING,
				Description: "Subscription plan.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Service type code.",
			},
			{
				Name:        "cloud_name",
				Type:        proto.ColumnType_STRING,
				Description: "Target cloud.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Service creation timestamp (ISO 8601).",
			},
			{
				Name:        "node_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of service nodes in the active plan.",
			},
			{
				Name:        "powered",
				Type:        proto.ColumnType_BOOL,
				Description: "Power-on the service (true) or power-off (false).",
			},
			{
				Name:        "project_vpc_id",
				Type:        proto.ColumnType_STRING,
				Description: "The Project VPC ID.",
				Transform:   transform.FromField("ProjectVPCID"),
			},
			{
				Name:        "termination_protection",
				Type:        proto.ColumnType_BOOL,
				Description: "Service is protected against termination and powering off.",
			},
			{
				Name:        "uri",
				Type:        proto.ColumnType_STRING,
				Description: "URI for connecting to the service (may be null).",
				Transform:   transform.FromField("URI"),
			},
			{
				Name:        "update_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Service last update timestamp (ISO 8601).",
			},
			{
				Name:        "acl",
				Type:        proto.ColumnType_JSON,
				Description: "List of Kafka ACL entries.",
				Transform:   transform.FromField("ACL"),
			},
			{
				Name:        "backups",
				Type:        proto.ColumnType_JSON,
				Description: "List of backups for the service.",
			},
			{
				Name:        "components",
				Type:        proto.ColumnType_JSON,
				Description: "Service component information objects.",
			},
			{
				Name:        "connection_info",
				Type:        proto.ColumnType_JSON,
				Description: "Service-specific connection information properties.",
			},
			{
				Name:        "connection_pools",
				Type:        proto.ColumnType_JSON,
				Description: "PostgreSQL PGBouncer connection pools.",
			},
			{
				Name:        "group_list",
				Type:        proto.ColumnType_JSON,
				Description: "List of service groups the service belongs to. This field is deprecated. It is always set to single element with value 'default'.",
			},
			{
				Name:        "integrations",
				Type:        proto.ColumnType_JSON,
				Description: "Integrations with other services.",
			},
			{
				Name:        "maintenance_window",
				Type:        proto.ColumnType_JSON,
				Description: "Automatic maintenance settings.",
			},
			{
				Name:        "metadata",
				Type:        proto.ColumnType_JSON,
				Description: "Service type specific metadata.",
			},
			{
				Name:        "node_states",
				Type:        proto.ColumnType_JSON,
				Description: "State of individual service nodes.",
			},
			{
				Name:        "uri_params",
				Type:        proto.ColumnType_JSON,
				Description: "Service URI parameterized into key-value pairs.",
				Transform:   transform.FromField("URIParams"),
			},
			{
				Name:        "user_config",
				Type:        proto.ColumnType_JSON,
				Description: "Service type-specific settings.",
			},
			{
				Name:        "users",
				Type:        proto.ColumnType_JSON,
				Description: "List of service users.",
			},
		},
	}
}

type AivenService struct {
	ProjectName string
	aiven.Service
}

func listServices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*aiven.Project)

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listServices", "connection_error", err)
		return nil, err
	}

	serviceList, err := conn.Services.List(project.Name)
	if err != nil {
		plugin.Logger(ctx).Error("listServices", "api_error", err)
		return nil, err
	}

	for _, service := range serviceList {
		d.StreamListItem(ctx, AivenService{project.Name, *service})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getService(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	project := d.EqualsQuals["project_name"].GetStringValue()
	name := d.EqualsQuals["name"].GetStringValue()

	// Check if project or name is empty.
	if project == "" || name == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getService", "connection_error", err)
		return nil, err
	}

	service, err := conn.Services.Get(project, name)
	if err != nil {
		plugin.Logger(ctx).Error("getService", "api_error", err)
		return nil, err
	}

	return AivenService{project, *service}, nil
}
