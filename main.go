package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	webServerPort := ":3000"

	startupWebServer(webServerPort)

}

func startupWebServer(webServerPort string) {

	// add handler functions here when they get written
	http.HandleFunc("/hello", helloHandler)

	log.Printf("Starting up webserver on port 3000...")
	if err := http.ListenAndServe(webServerPort, nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}
