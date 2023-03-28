package aiven

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAivenServiceIntegrationEndpoint(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_service_integration_endpoint",
		Description: "Retrieve information about your service integration endpoints.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listServiceIntegrationEndpoints,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"endpoint_id", "project_name"}),
			Hydrate:    getServiceIntegrationEndpoint,
		},
		Columns: []*plugin.Column{
			{
				Name:        "endpoint_id",
				Type:        proto.ColumnType_STRING,
				Description: "The integration endpoint ID.",
				Transform:   transform.FromField("EndpointID"),
			},
			{
				Name:        "endpoint_name",
				Type:        proto.ColumnType_STRING,
				Description: "The integration endpoint name.",
			},
			{
				Name:        "project_name",
				Type:        proto.ColumnType_STRING,
				Description: "The project name.",
			},
			{
				Name:        "endpoint_type",
				Type:        proto.ColumnType_STRING,
				Description: "Integration endpoint type.",
			},
			{
				Name:        "endpoint_config",
				Type:        proto.ColumnType_JSON,
				Description: "Service integration endpoint backend settings.",
			},
			{
				Name:        "user_config",
				Type:        proto.ColumnType_JSON,
				Description: "Service integration endpoint settings.",
			},
		},
	}
}

type IntegrationEndpoint struct {
	ProjectName string
	aiven.ServiceIntegrationEndpoint
}

func listServiceIntegrationEndpoints(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*aiven.Project)

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listServiceIntegrationEndpoints", "connection_error", err)
		return nil, err
	}

	integrationEndpoints, err := conn.ServiceIntegrationEndpoints.List(project.Name)
	if err != nil {
		plugin.Logger(ctx).Error("listServiceIntegrationEndpoints", "api_error", err)
		return nil, err
	}

	for _, integrationEndpoint := range integrationEndpoints {
		d.StreamListItem(ctx, IntegrationEndpoint{project.Name, *integrationEndpoint})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getServiceIntegrationEndpoint(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	project := d.EqualsQuals["project_name"].GetStringValue()
	id := d.EqualsQuals["endpoint_id"].GetStringValue()

	// Check if project or id is empty.
	if project == "" || id == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getServiceIntegrationEndpoint", "connection_error", err)
		return nil, err
	}

	integrationEndpoint, err := conn.ServiceIntegrationEndpoints.Get(project, id)
	if err != nil {
		plugin.Logger(ctx).Error("getServiceIntegrationEndpoint", "api_error", err)
		return nil, err
	}

	return IntegrationEndpoint{project, *integrationEndpoint}, nil
}
