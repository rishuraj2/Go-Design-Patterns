package main

import (
	trafficlight "trafficlightcontroller/internal/trafficLight"
	"trafficlightcontroller/internal/trafficLight/state"
)

func main() {
	trafficLight := trafficlight.NewTrafficLight(state.NewRedState())
	trafficLight.Change()
	trafficLight.Change()
	trafficLight.Change()
	trafficLight.Change()
}
