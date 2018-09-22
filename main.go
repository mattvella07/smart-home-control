package main

import (
	"log"
	"net/http"

	"github.com/mattvella07/smart-home-control/api"
)

func healthCheck(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte("Success"))
}

func createServer() {
	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/api/hue/getLights", api.GetLights)
	http.HandleFunc("/api/hue/turnOnLight/", api.TurnOnLight)
	http.HandleFunc("/api/hue/turnOffLight/", api.TurnOffLight)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	createServer()
}
