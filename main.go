package main

import (
	"fmt"
	"io/ioutil"
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
	http.HandleFunc("/stats", statsHandler)

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

func getWarState() []byte {
	resp, err := http.Get("https://war-service-live-2.foxholeservices.com/api/worldconquest/war")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/stats" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, string(getWarState()))
}
