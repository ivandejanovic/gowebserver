package main

import (
	"fmt"
	"goWebServerBasic/router"
	"net/http"
)

func main() {
	fmt.Println("Started")
	var addr string = ":8000"
	var handler http.Handler = router.CreateRouteHandler()

	// Bind to a port and pass route handler in
	server := http.Server{Addr: addr, Handler: handler}

	fmt.Println("Fire")
	server.ListenAndServe()
}
