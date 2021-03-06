package main

import (
	"github.com/ivandejanovic/gowebserver/router"
	"net/http"
)

func main() {
	var addr string = ":8080"
	var handler http.Handler = router.CreateRouteHandler()

	// Bind to a port and pass route handler in
	server := http.Server{Addr: addr, Handler: handler}

	server.ListenAndServe()
}
