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
	http.Handle("/healthcheck", m.MethodChecker(http.HandlerFunc(healthCheck)))
	http.Handle("/api/hue/getLights", m.MethodChecker(http.HandlerFunc(api.GetLights)))
	http.Handle("/api/hue/getLight/", m.MethodChecker(http.HandlerFunc(api.GetLight)))
	http.Handle("/api/nest/getThermostats", m.MethodChecker(http.HandlerFunc(api.GetThermostats)))
	http.Handle("/api/nest/getThermostat/", m.MethodChecker(http.HandlerFunc(api.GetThermostat)))

	m.Allowed = []string{"POST"}
	http.Handle("/api/hue/turnOnLight/", m.MethodChecker(http.HandlerFunc(api.TurnOnLight)))
	http.Handle("/api/hue/turnOffLight/", m.MethodChecker(http.HandlerFunc(api.TurnOffLight)))
	http.Handle("/api/nest/setTargetTemperature/", m.MethodChecker(http.HandlerFunc(api.SetTargetTemperature)))
	http.Handle("/api/nest/setTargetHighLowTemperature/", m.MethodChecker(http.HandlerFunc(api.SetTargetHighLowTemperature)))
	http.Handle("/api/nest/setHVACMode/", m.MethodChecker(http.HandlerFunc(api.SetHVACMode)))
	http.Handle("/api/nest/setThermostatLabel/", m.MethodChecker(http.HandlerFunc(api.SetThermostatLabel)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	createServer()
}
