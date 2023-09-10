package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title    string
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("pages/index.gohtml", "templates/base.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "main", PageData{"example 1"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("pages/other.gohtml", "templates/base.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "main", PageData{"example 2"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}