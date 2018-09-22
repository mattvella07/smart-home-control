package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/mattvella07/hue"
)

func GetLights(rw http.ResponseWriter, r *http.Request) {
	var h hue.Connection

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

func TurnOnLight(rw http.ResponseWriter, r *http.Request) {
	reqURL := r.URL.String()
	if strings.Index(reqURL, "?") > -1 {
		reqURL = reqURL[0:strings.Index(reqURL, "?")]
	}

	params := strings.Split(strings.Replace(reqURL, "/api/hue/turnOnLight/", "", 1), "/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) > 0 {
		if light, err := strconv.Atoi(params[0]); err == nil {
			var h hue.Connection

			err = h.TurnOnLight(light)
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

func TurnOffLight(rw http.ResponseWriter, r *http.Request) {
	reqURL := r.URL.String()
	if strings.Index(reqURL, "?") > -1 {
		reqURL = reqURL[0:strings.Index(reqURL, "?")]
	}

	params := strings.Split(strings.Replace(reqURL, "/api/hue/turnOffLight/", "", 1), "/")

	rw.Header().Set("Content-Type", "text/html")

	if len(params) > 0 {
		if light, err := strconv.Atoi(params[0]); err == nil {
			var h hue.Connection

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
