package aiven

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAivenServiceIntegration(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_service_integration",
		Description: "Retrieve information about your service integrations.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"source_project", "source_service"}),
			Hydrate:    listServiceIntegrations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"source_project", "service_integration_id"}),
			Hydrate:    getServiceIntegration,
		},
		Columns: []*plugin.Column{
			{
				Name:        "service_integration_id",
				Type:        proto.ColumnType_STRING,
				Description: "The service integration ID.",
				Transform:   transform.FromField("ServiceIntegrationID"),
			},
			{
				Name:        "source_project",
				Type:        proto.ColumnType_STRING,
				Description: "The source project.",
			},
			{
				Name:        "source_service",
				Type:        proto.ColumnType_STRING,
				Description: "The source service.",
			},
			{
				Name:        "active",
				Type:        proto.ColumnType_BOOL,
				Description: "True when integration is active.",
			},
			{
				Name:        "enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "True when integration is enabled.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description of the integration.",
			},
			{
				Name:        "destination_endpoint_id",
				Type:        proto.ColumnType_STRING,
				Description: "The destination endpoint ID.",
				Transform:   transform.FromField("DestinationEndpointID"),
			},
			{
				Name:        "destination_endpoint_name",
				Type:        proto.ColumnType_STRING,
				Description: "The destination endpoint name.",
			},
			{
				Name:        "destination_project",
				Type:        proto.ColumnType_STRING,
				Description: "The destination project.",
			},
			{
				Name:        "destination_service",
				Type:        proto.ColumnType_STRING,
				Description: "The destination service.",
			},
			{
				Name:        "destination_service_type",
				Type:        proto.ColumnType_STRING,
				Description: "The destination service type code.",
			},
			{
				Name:        "integration_type",
				Type:        proto.ColumnType_STRING,
				Description: "Type of the integration.",
			},
			{
				Name:        "source_endpoint_id",
				Type:        proto.ColumnType_STRING,
				Description: "The source endpoint ID.",
				Transform:   transform.FromField("SourceEndpointID"),
			},
			{
				Name:        "source_endpoint_name",
				Type:        proto.ColumnType_STRING,
				Description: "The source endpoint name.",
			},
			{
				Name:        "source_service_type",
				Type:        proto.ColumnType_STRING,
				Description: "The source service type code.",
			},
			{
				Name:        "integration_status",
				Type:        proto.ColumnType_JSON,
				Description: "Integration status.",
			},
			{
				Name:        "user_config",
				Type:        proto.ColumnType_JSON,
				Description: "Service integration settings.",
			},
		},
	}
}

func listServiceIntegrations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	project := d.EqualsQuals["source_project"].GetStringValue()
	service := d.EqualsQuals["source_service"].GetStringValue()

	// Check if project or service is empty.
	if project == "" || service == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listServiceIntegrations", "connection_error", err)
		return nil, err
	}

	serviceintegrations, err := conn.ServiceIntegrations.List(project, service)
	if err != nil {
		plugin.Logger(ctx).Error("listServiceIntegrations", "api_error", err)
		return nil, err
	}

	for _, integration := range serviceintegrations {
		d.StreamListItem(ctx, integration)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getServiceIntegration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	project := d.EqualsQuals["source_project"].GetStringValue()
	integrationID := d.EqualsQuals["service_integration_id"].GetStringValue()

	// Check if project or integrationID is empty.
	if project == "" || integrationID == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getServiceIntegration", "connection_error", err)
		return nil, err
	}

	integration, err := conn.ServiceIntegrations.Get(project, integrationID)
	if err != nil {
		plugin.Logger(ctx).Error("getServiceIntegration", "api_error", err)
		return nil, err
	}

	return integration, nil
}
