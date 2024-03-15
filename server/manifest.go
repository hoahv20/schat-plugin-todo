// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.schat.plugin-todo-next",
  "name": "Todo",
  "description": "This plugin makes it easy to keep track of Todo issues and get daily reminders.",
  "homepage_url": "https://github.com/mattermost/mattermost-plugin-todo",
  "support_url": "https://github.com/mattermost/mattermost-plugin-todo/issues",
  "release_notes_url": "https://github.com/mattermost/mattermost-plugin-todo/releases/tag/v0.7.1",
  "version": "0.7.1",
  "min_server_version": "6.5.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "hide_team_sidebar",
        "display_name": "Hide team sidebar buttons:",
        "type": "bool",
        "help_text": "When true, the buttons in the team sidebar on the left toolbar will be hidden.",
        "placeholder": "",
        "default": null
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
