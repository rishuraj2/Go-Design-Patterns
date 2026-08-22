package state

import (
	"fmt"
	"time"
	trafficlight "trafficlightcontroller/internal/trafficLight"
)

type GreenState struct {
	durationSeconds int
}

func NewGreenState() *GreenState {
	return &GreenState{
		durationSeconds: 5,
	}
}

func (this *GreenState) Change(context *trafficlight.TrafficLight) {
	fmt.Println("GREEN light - Go")
	time.Sleep(time.Duration(this.durationSeconds)*time.Second)
	context.SetLightState(NewYellowState())
}
