package main

import (
	"flag"
	"net/http"

	"github.com/haxii/daemon"

	"sbdiplom/internal/data"
	"sbdiplom/internal/infrastructure"
)

var (
	dport = flag.Int("p", 8282, "server port")
	_     = flag.String("s", daemon.UsageDefaultName, daemon.UsageMessage)
)

func main() {
	daemon.Make("-s", "httpdaemon", "simple http daemon service").Run(serve)
}

func serve() {
	flag.Parse()
	http.HandleFunc("/api", data.HandleConnection)
	infrastructure.Run(dport)
}
