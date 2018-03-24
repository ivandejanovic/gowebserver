package router

import (
	"github.com/gorilla/mux"
	"github.com/ivandejanovic/gowebserver/routes"
	"net/http"
)

func setStaticRoutesHandler(router *mux.Router) {
	// Routes consist of a path and a handler function.
	path := http.Dir("src/github.com/ivandejanovic/gowebserver/static/")
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
