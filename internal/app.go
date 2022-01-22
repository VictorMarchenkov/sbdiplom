package infrastructure

import (
	"fmt"
	"net/http"
)

func Run(port *int) {
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
