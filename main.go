package main

import (
	"log"
	"net/http"

	"github.com/mattvella07/smart-home-control/phillips"
)

func healthCheck(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte("Success"))
}

func createServer() {
	var hue phillips.Hue

	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/api/hue/getLights", hue.GetLights)
	http.HandleFunc("/api/hue/turnOnLight/", hue.TurnOnLight)
	http.HandleFunc("/api/hue/turnOffLight/", hue.TurnOffLight)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	createServer()
}
