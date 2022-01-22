package main

import (
	"flag"
	"github.com/haxii/daemon"

	"sbdiplom/internal"
)

var (
	port = flag.Int("p", 8282, "server port")
	_    = flag.String("s", daemon.UsageDefaultName, daemon.UsageMessage)
)

func main() {
	flag.Parse()
	//	d := daemon.Make("-s", "diploma", "simple http daemon service")
	//	d.Run(func() {
	internal.Run(port)
	//	})
}
