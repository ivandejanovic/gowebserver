package router

import (
	"github.com/gorilla/mux"
	"github.com/ivandejanovic/gowebserver/routes"
	"net/http"
)

func setStaticRoutesHandler(router *mux.Router) {
	// Routes consist of a path and a handler function.
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
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
