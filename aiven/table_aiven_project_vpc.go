package aiven

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAivenProjectVpc(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_project_vpc",
		Description: "Retrieve information about your project VPCs.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listProjectVpcs,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project_vpc_id", "project_name"}),
			Hydrate:    getProjectVpc,
		},
		Columns: []*plugin.Column{
			{
				Name:        "project_vpc_id",
				Type:        proto.ColumnType_STRING,
				Description: "The project VPC ID.",
				Transform:   transform.FromField("ProjectVPCID"),
			},
			{
				Name:        "project_name",
				Type:        proto.ColumnType_STRING,
				Description: "The project name.",
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "The Project VPC state.",
			},
			{
				Name:        "cloud_name",
				Type:        proto.ColumnType_STRING,
				Description: "The target cloud.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "VPC creation timestamp.",
			},
			{
				Name:        "network_cidr",
				Type:        proto.ColumnType_CIDR,
				Description: "IPv4 network range CIDR.",
				Transform:   transform.FromField("NetworkCIDR"),
			},
			{
				Name:        "update_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp of last change to VPC.",
			},
			{
				Name:        "peering_connections",
				Type:        proto.ColumnType_JSON,
				Description: "List of peering connections.",
			},
		},
	}
}

type ProjectVpc struct {
	ProjectName string
	aiven.VPC
}

func listProjectVpcs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*aiven.Project)

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listProjectVpcs", "connection_error", err)
		return nil, err
	}

	vpcList, err := conn.VPCs.List(project.Name)
	if err != nil {
		plugin.Logger(ctx).Error("listProjectVpcs", "api_error", err)
		return nil, err
	}

	for _, vpc := range vpcList {
		d.StreamListItem(ctx, ProjectVpc{project.Name, *vpc})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getProjectVpc(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	project := d.EqualsQuals["project_name"].GetStringValue()
	id := d.EqualsQuals["project_vpc_id"].GetStringValue()

	// Check if project or id is empty.
	if project == "" || id == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getProjectVpc", "connection_error", err)
		return nil, err
	}

	vpc, err := conn.VPCs.Get(project, id)
	if err != nil {
		plugin.Logger(ctx).Error("getProjectVpc", "api_error", err)
		return nil, err
	}

	return ProjectVpc{project, *vpc}, nil
}
