package main

import (
	"flag"
	"github.com/haxii/daemon"

	"sbdaemon/internal/app/v2"
)

var (
	port = flag.Int("p", 8282, "server port")
	_    = flag.String("s", daemon.UsageDefaultName, daemon.UsageMessage)
)

func main() {
	flag.Parse()
	d := daemon.Make("-s", "main", "simple http daemon service")
	d.Run(func() {
		app.Run(*port)
	})
}
