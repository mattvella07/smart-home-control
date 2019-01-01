package main

import (
	"log"
	"net/http"

	"github.com/mattvella07/smart-home-control/api"
	"github.com/mattvella07/smart-home-control/middleware"
)

func healthCheck(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte("Success"))
}

func createServer() {
	m := middleware.Method{}

	m.Allowed = []string{"GET"}
	http.Handle("/", m.MethodChecker(http.HandlerFunc(healthCheck)))
	http.Handle("/api/hue/getLights", m.MethodChecker(http.HandlerFunc(api.GetLights)))

	m.Allowed = []string{"POST"}
	http.Handle("/api/hue/turnOnLight/", m.MethodChecker(http.HandlerFunc(api.TurnOnLight)))
	http.Handle("/api/hue/turnOffLight/", m.MethodChecker(http.HandlerFunc(api.TurnOffLight)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	createServer()
}
