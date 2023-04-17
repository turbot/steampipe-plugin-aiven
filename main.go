package main

import (
	"github.com/turbot/steampipe-plugin-aiven/aiven"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: aiven.Plugin})
}
