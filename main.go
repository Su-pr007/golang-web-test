package main

import (
	"os"
	"log"
	"html/template"
	"net/http"
	"github.com/joho/godotenv"
)

type PageData struct {
	Title    string
}

// init is invoked before main()
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	appPort, exists := os.LookupEnv("APP_PORT")

	if !exists {
		log.Print("Port is not set")
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":"+appPort, nil)
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