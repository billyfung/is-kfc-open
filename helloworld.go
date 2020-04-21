package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Model of stuff to render a page
type restaurant struct {
	Name string
	Open bool
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("received a request.")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// t := time.Now()

	tmpl.Execute(w, restaurant{"KFC", false})
}

func main() {
	log.Print("started main()")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
