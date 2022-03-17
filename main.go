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
	internal.Run(port)
}
