package main

import (
	"net/http"
	"router"
)

func main() {
	var addr string = ":8000"
	var handler http.Handler = router.CreateRouteHandler()

	// Bind to a port and pass route handler in
	server := http.Server{Addr: addr, Handler: handler}

	server.ListenAndServe()
}
