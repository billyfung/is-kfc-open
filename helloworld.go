package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// Model of stuff to render a page
type restaurant struct {
	Name string
	Open bool
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("received a request.")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	//akl, _ := time.LoadLocation("Pacific/Auckland")
	// convert to WIB
	kfc := restaurant{"KFC", false}
	t := time.Now().Unix()
	if t >= 1588021200 {
		kfc.Open = true
	} else {
		kfc.Open = false
	}
	tmpl.Execute(w, kfc)
}

func main() {
	log.Print("started server()")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
