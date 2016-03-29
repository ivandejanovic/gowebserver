package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
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

func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("tmpl/home.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Parsing home template failed: " + err.Error()))
	}
	data := struct {
		Title string
	}{
		Title: "Go Web Server Basic Home Page",
	}
	tmpl.Execute(w, data)
}
