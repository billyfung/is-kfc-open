package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("received a request.")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}
	t := time.Now().UTC()

	fmt.Fprintf(w, "Hello %s!, UTC time is %s\n", target, t)
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
