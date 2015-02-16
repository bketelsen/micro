package main

import (
	"github.com/asim/go-micro/cmd"
	"github.com/asim/go-micro/server"
	"github.com/bketelsen/micro/server/handler"
	log "github.com/golang/glog"
)

func main() {
	// optionally setup command line usage
	cmd.Init()

	server.Name = "sample.math"

	// Initialise Server
	server.Init()

	// Register Handlers
	server.Register(
		server.NewReceiver(
			new(handler.Math),
		),
	)

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
