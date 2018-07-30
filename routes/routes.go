package routes

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ivandejanovic/gowebserver/service"
)

func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["key"]
	fmt.Println(value)
	response, err := http.Get("http://golang.org/")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Request to golang site failed"))
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to read body content"))
	}
	w.Write([]byte(contents))
}

func Ajax2Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["key"]
	fmt.Println(value)
	response, err := service.Method(value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Request to service failed"))
	}

	w.Write([]byte(response))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var path string
	static, present := os.LookupEnv("STATIC")
	if present {
		path = static + "/tmpl/home.html"
	} else {
		path = "src/github.com/ivandejanovic/gowebserver/tmpl/home.html"
	}
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Parsing home template failed: " + err.Error()))
	}
	data := struct {
		Title string
	}{
		Title: "Go Web Server Home Page",
	}
	tmpl.Execute(w, data)
}
