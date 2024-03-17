package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/tucnak/steampipe-plugin-slack/slack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: slack.Plugin})
}
