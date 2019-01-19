package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mattvella07/nest"
)

// GetThermostats returns all Nest thermostats
// Route: GET /api/nest/getThermostats
// Parms: None
func GetThermostats(rw http.ResponseWriter, r *http.Request) {
	n := nest.Connection{
		AccessToken: os.Getenv("nestAccessToken"),
	}

	thermostats, err := n.GetThermostats()
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Header().Set("content-type", "application/json")
	json.NewEncoder(rw).Encode(thermostats)
}

// GetThermostat returns the specified Nest thermostat
// Route: GET /api/nest/getThermostat/{thermostatID}
// Parms: None
func GetThermostat(rw http.ResponseWriter, r *http.Request) {
	params := extractPathParams(r.URL.String(), "/api/nest/getThermostat/")

	if len(params) == 0 {
		rw.WriteHeader(500)
		rw.Write([]byte("Must include the thermostat ID: /api/nest/getThermostat/{thermostatID}"))
		return
	}

	thermostatID := params[0]

	n := nest.Connection{
		AccessToken: os.Getenv("nestAccessToken"),
	}

	thermostat, err := n.GetThermostat(thermostatID)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Header().Set("content-type", "application/json")
	json.NewEncoder(rw).Encode(thermostat)
}

// SetTargetTemperature sets the target temperature on the specified Nest thermostat
// Route: POST /api/nest/setTargetTemperature/{thermostatID}
// Params:
//     temp - The target temperature for the thermostat to be set to
// * This should only be called if the HVAC Mode is not set to heat-cool *
func SetTargetTemperature(rw http.ResponseWriter, r *http.Request) {
	params := extractPathParams(r.URL.String(), "/api/nest/setTargetTemperature/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) == 0 {
		rw.WriteHeader(500)
		rw.Write([]byte("Must include the thermostat ID: /api/nest/setTargetTemperature/{thermostatID}"))
		return
	}

	thermostatID := params[0]

	temp := r.URL.Query().Get("temp")
	if strings.Trim(temp, " ") == "" {
		rw.WriteHeader(500)
		rw.Write([]byte("Query parameter temp must be provided"))
		return
	}

	n := nest.Connection{
		AccessToken: os.Getenv("nestAccessToken"),
	}

	scale, err := n.GetTemperatureScale(thermostatID)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	switch scale {
	case "F":
		tempF, err := strconv.Atoi(temp)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte("Invalid temperature value"))
			return
		}

		err = n.SetTargetTemperatureF(thermostatID, tempF)
	case "C":
		tempC, err := strconv.ParseFloat(temp, 32)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte("Invalid temperature value"))
			return
		}

		err = n.SetTargetTemperatureC(thermostatID, float32(tempC))
	default:
		rw.WriteHeader(500)
		rw.Write([]byte("Error while getting temperature scale"))
		return
	}

	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write([]byte(fmt.Sprintf("Target temperature set to %s", temp)))
}

// SetTargetHighLowTemperature sets the target high and low temperature on the specified Nest thermostat
// Route: POST /api/nest/setTargetHighLowTemperature/{thermostatID}
// Params:
//     high - The target high temperature for the thermostat to be set to
//     low - The target low temperature for the thermostat to be set to
// * This should only be called if the HVAC Mode is set to heat-cool *
func SetTargetHighLowTemperature(rw http.ResponseWriter, r *http.Request) {
	params := extractPathParams(r.URL.String(), "/api/nest/setTargetHighLowTemperature/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) == 0 {
		rw.WriteHeader(500)
		rw.Write([]byte("Must include the thermostat ID: /api/nest/setTargetHighLowTemperature/{thermostatID}"))
	}

	thermostatID := params[0]

	high := r.URL.Query().Get("high")
	if strings.Trim(high, " ") == "" {
		rw.WriteHeader(500)
		rw.Write([]byte("Query parameter high must be provided"))
	}

	low := r.URL.Query().Get("low")
	if strings.Trim(low, " ") == "" {
		rw.WriteHeader(500)
		rw.Write([]byte("Query parameter low must be provided"))
	}

	n := nest.Connection{
		AccessToken: os.Getenv("nestAccessToken"),
	}

	scale, err := n.GetTemperatureScale(thermostatID)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	switch scale {
	case "F":
		highF, err := strconv.Atoi(high)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte("Invalid high temperature value"))
			return
		}

		lowF, err := strconv.Atoi(low)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte("Invalid low temperature value"))
			return
		}

		err = n.SetTargetHighLowTemperatureF(thermostatID, highF, lowF)
	case "C":
		highC, err := strconv.ParseFloat(high, 32)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte("Invalid high temperature value"))
			return
		}

		lowC, err := strconv.ParseFloat(low, 32)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte("Invalid low temperature value"))
			return
		}

		err = n.SetTargetHighLowTemperatureC(thermostatID, float32(highC), float32(lowC))
	default:
		rw.WriteHeader(500)
		rw.Write([]byte("Error while getting temperature scale"))
		return
	}

	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write([]byte(fmt.Sprintf("Target high temperature set to %s, and target low temperature set to %s", high, low)))
}

// SetHVACMode sets the HVAC mode on the specified Nest thermostat
// Route: POST /api/nest/setHVACMode/{thermostatID}
// Params:
//     mode - The HVAC mode. Valid values: ["heat", "cool", "heat-cool", "eco", "off"]
func SetHVACMode(rw http.ResponseWriter, r *http.Request) {
	params := extractPathParams(r.URL.String(), "/api/nest/setHVACMode/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) == 0 {
		rw.WriteHeader(500)
		rw.Write([]byte("Must include the thermostat ID: /api/nest/setHVACMode/{thermostatID}"))
	}

	thermostatID := params[0]
	mode := r.URL.Query().Get("mode")

	n := nest.Connection{
		AccessToken: os.Getenv("nestAccessToken"),
	}

	err := n.SetHVACMode(thermostatID, mode)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write([]byte(fmt.Sprintf("HVAC Mode set to %s", mode)))
}

// SetThermostatLabel sets the label on the specified Nest thermostat
// Route: POST /api/nest/setLabel/{thermostatID}
// Params:
//     label - The text that identifies the thermostat
func SetThermostatLabel(rw http.ResponseWriter, r *http.Request) {
	params := extractPathParams(r.URL.String(), "/api/nest/setLabel/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) == 0 {
		rw.WriteHeader(500)
		rw.Write([]byte("Must include the thermostat ID: /api/nest/setLabel/{thermostatID}"))
	}

	thermostatID := params[0]
	label := r.URL.Query().Get("label")

	n := nest.Connection{
		AccessToken: os.Getenv("nestAccessToken"),
	}

	err := n.SetThermostatLabel(thermostatID, label)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write([]byte(fmt.Sprintf("Label set to %s", label)))
}
