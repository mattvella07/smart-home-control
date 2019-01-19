package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/mattvella07/hue"
)

// GetLights returns all Phillips Hue lights
// Route: GET /api/hue/getLights
// Parms: None
func GetLights(rw http.ResponseWriter, r *http.Request) {
	h := hue.Connection{
		UserID: os.Getenv("hueUserID"),
	}

	lights, err := h.GetLights()
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(lights)
}

// GetLight returns the specified Phillips hue light
// Route: GET /api/hue/getLight/{lightID}
// Parms: None
func GetLight(rw http.ResponseWriter, r *http.Request) {
	params := extractPathParams(r.URL.String(), "/api/hue/getLight/")

	if len(params) == 0 {
		rw.WriteHeader(500)
		rw.Write([]byte("Must include the light ID: /api/hue/getLight/{lightID}"))
		return
	}

	lightID, err := strconv.Atoi(params[0])
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("Invalid light ID"))
		return
	}

	h := hue.Connection{
		UserID: os.Getenv("hueUserID"),
	}

	light, err := h.GetLight(lightID)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Header().Set("content-type", "application/json")
	json.NewEncoder(rw).Encode(light)
}

// TurnOnLight turns on the specified Phillips Hue light
// and sets the color if the parameters are supplied
// Route: POST /api/hue/turnOnLight/{lightID}
// Parms:
//     x -
//     y -
//     bri -
//     hue -
//     sat -
// * Query parameters are optional, but all must be provided if any are to be used *
func TurnOnLight(rw http.ResponseWriter, r *http.Request) {
	params := extractPathParams(r.URL.String(), "/api/hue/turnOnLight/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) == 0 {
		rw.WriteHeader(500)
		rw.Write([]byte("Must include the light ID: /api/hue/turnOnLight/{lightID}"))
		return
	}

	lightID, err := strconv.Atoi(params[0])
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("Invalid light ID"))
		return
	}

	h := hue.Connection{
		UserID: os.Getenv("hueUserID"),
	}

	x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 32)
	if err != nil {
		x = 0.0
	}

	y, err := strconv.ParseFloat(r.URL.Query().Get("y"), 32)
	if err != nil {
		y = 0.0
	}

	bri, err := strconv.Atoi(r.URL.Query().Get("bri"))
	if err != nil {
		bri = 0
	}

	hue, err := strconv.Atoi(r.URL.Query().Get("hue"))
	if err != nil {
		hue = 0
	}

	sat, err := strconv.Atoi(r.URL.Query().Get("sat"))
	if err != nil {
		sat = 0
	}

	if x == 0.0 && y == 0.0 && bri == 0 && hue == 0 && sat == 0 {
		err = h.TurnOnLight(lightID)
	} else {
		err = h.TurnOnLightWithColor(lightID, float32(x), float32(y), bri, hue, sat)
	}

	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write([]byte(fmt.Sprintf("Light %d turned on", lightID)))
}

// TurnOffLight turns off the specified Phillips Hue light
// Route: POST /api/hue/turnOffLight/{lightID}
// Parms: None
func TurnOffLight(rw http.ResponseWriter, r *http.Request) {
	params := extractPathParams(r.URL.String(), "/api/hue/turnOffLight/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) == 0 {
		rw.WriteHeader(500)
		rw.Write([]byte("Must include the light ID: /api/hue/turnOffLight/{lightID}"))
		return
	}

	lightID, err := strconv.Atoi(params[0])
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("Invalid light ID"))
		return
	}

	h := hue.Connection{
		UserID: os.Getenv("hueUserID"),
	}

	err = h.TurnOffLight(lightID)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write([]byte(fmt.Sprintf("Light %d turned off", lightID)))
}
