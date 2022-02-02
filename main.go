package main

import (
	c "search-engine/config"
	"search-engine/server"

	"github.com/golang/glog"
)

func main() {
	// We build app with custom params
	app, err := server.NewApp(c.Config.Debug)
	if err != nil {
		glog.Fatal(err)
	}

	// Run server at specific port
	if err := app.Run(c.Config.Port); err != nil {
		glog.Fatal(err)
	}
}
