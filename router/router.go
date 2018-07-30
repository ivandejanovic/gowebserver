package router

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ivandejanovic/gowebserver/routes"
)

func setStaticRoutesHandler(router *mux.Router) {
	// Routes consist of a path and a handler function.
	var path http.Dir
	static, present := os.LookupEnv("STATIC")
	if present {
		path = http.Dir(static + "/static")
	} else {
		path = http.Dir("src/github.com/ivandejanovic/gowebserver/static/")
	}
	sh := http.FileServer(path)
	s := http.StripPrefix("/static/", sh)
	router.PathPrefix("/static/").Handler(s)
}

func setRouteHandler(router *mux.Router) {
	router.HandleFunc("/ajax/{key}", routes.AjaxHandler)
	router.HandleFunc("/ajax2/{key}", routes.Ajax2Handler)
	router.HandleFunc("/", routes.RootHandler)
}

func CreateRouteHandler() http.Handler {
	router := mux.NewRouter()

	setStaticRoutesHandler(router)
	setRouteHandler(router)

	return router
}
