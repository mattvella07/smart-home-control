package main

import (
	"fmt"
	"log"

	"github.com/mattvella07/smart-home-control/phillips"
)

func main() {
	var hue phillips.Hue

	err := hue.GetBridgeIPAddress()
	if err != nil {
		log.Fatalf("GetBridgeIPAddress Error: %s", err)
	}

	err = hue.GetUserID()
	if err != nil {
		log.Fatalf("GetUserID Error: %s", err)
	}

	hue.GetBaseURL()

	err = hue.GetLights()
	if err != nil {
		log.Fatalf("GetLights Error: %s", err)
	}

	fmt.Println("Phillips Hue:")
	for _, light := range hue.Lights {
		fmt.Println(light.Name)
	}

	err = hue.ChangeLightState(4, "on", "false")
	if err != nil {
		log.Fatalf("ChangeLightState Error: %s", err)
	}
}
