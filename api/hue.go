package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mattvella07/hue"
)

// GetLights returns all Phillips Hue lights
func GetLights(rw http.ResponseWriter, r *http.Request) {
	h := hue.Connection{
		UserID: os.Getenv("hueUserID"),
	}

	lights, err := h.GetLights()
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		rw.WriteHeader(500)
		rw.Write([]byte(fmt.Sprintf("ERROR: %s", err)))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(lights)
}

// TurnOnLight turns on the specified Phillips Hue light
// and sets the color if the parameters are supplied
func TurnOnLight(rw http.ResponseWriter, r *http.Request) {
	reqURL := r.URL.String()
	if strings.Index(reqURL, "?") > -1 {
		reqURL = reqURL[0:strings.Index(reqURL, "?")]
	}

	params := strings.Split(strings.Replace(reqURL, "/api/hue/turnOnLight/", "", 1), "/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) > 0 {
		if light, err := strconv.Atoi(params[0]); err == nil {
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
				err = h.TurnOnLight(light)
			} else {
				err = h.TurnOnLightWithColor(light, float32(x), float32(y), bri, hue, sat)
			}

			if err != nil {
				rw.WriteHeader(500)
				rw.Write([]byte(err.Error()))
				return
			}

			rw.Write([]byte("Success"))
			return
		}
	}

	rw.WriteHeader(500)
	rw.Write([]byte("Error turning on light"))
}

// TurnOffLight turns off the specified Phillips Hue light
func TurnOffLight(rw http.ResponseWriter, r *http.Request) {
	reqURL := r.URL.String()
	if strings.Index(reqURL, "?") > -1 {
		reqURL = reqURL[0:strings.Index(reqURL, "?")]
	}

	params := strings.Split(strings.Replace(reqURL, "/api/hue/turnOffLight/", "", 1), "/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) > 0 {
		if light, err := strconv.Atoi(params[0]); err == nil {
			h := hue.Connection{
				UserID: os.Getenv("hueUserID"),
			}

			err = h.TurnOffLight(light)
			if err != nil {
				rw.WriteHeader(500)
				rw.Write([]byte(err.Error()))
				return
			}

			rw.Write([]byte("Success"))
			return
		}
	}

	rw.WriteHeader(500)
	rw.Write([]byte("Error turning off light"))
}
