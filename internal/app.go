package internal

import (
	"fmt"
	"net/http"
)

func Run(port *int) {
	var cfgCSV Config
	var resultT ResultT
	//r := mux.NewRouter()
	resultT.HandlerB(cfgCSV)
	http.HandleFunc("/api", resultT.HandlerT)

	fs := http.FileServer(http.Dir("internal/generator"))
	http.Handle("/", fs)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
