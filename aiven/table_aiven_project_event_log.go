package aiven

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAivenProjectEventLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_project_event_log",
		Description: "Retrieve information about your project event logs.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listProjectEventLogs,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "project_name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "actor",
				Type:        proto.ColumnType_STRING,
				Description: "Initiator of the event.",
			},
			{
				Name:        "event_desc",
				Type:        proto.ColumnType_STRING,
				Description: "The event description.",
			},
			{
				Name:        "event_type",
				Type:        proto.ColumnType_STRING,
				Description: "Event type identifier.",
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
				Name:        "time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp in ISO 8601 format, always in UTC.",
			},
		},
	}
}

type ProjectEventLog struct {
	ProjectName string
	aiven.ProjectEvent
}

func listProjectEventLogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*aiven.Project)
	project_name := d.EqualsQuals["project_name"].GetStringValue()

	// check if the provided project_name is not matching with the parentHydrate
	if project_name != "" && project_name != project.Name {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_project_event_log.listProjectEventLogs", "connection_error", err)
		return nil, err
	}

	eventLogs, err := conn.Projects.GetEventLog(project.Name)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_project_event_log.listProjectEventLogs", "api_error", err)
		return nil, err
	}

	for _, eventLog := range eventLogs {
		d.StreamListItem(ctx, ProjectEventLog{project.Name, *eventLog})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
