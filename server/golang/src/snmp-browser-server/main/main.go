package main

import (
	"golang/src/snmp-browser-server/main/app"
	"golang/src/snmp-browser-server/main/cli"
	"golang/src/snmp-browser-server/main/http"
)

func main() {
	appOptions := app.NewOptions()
	appOptions.FillFromCommandLine()

	if appOptions.Mode == app.HttpAppMode {
		http.Handle(appOptions)
	} else {
		cli.Handle(appOptions)
	}
}
