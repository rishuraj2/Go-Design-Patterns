package state

import (
	"fmt"
	"time"
	trafficlight "trafficlightcontroller/internal/trafficLight"
)

type YellowState struct {
	durationSeconds int
}

func NewYellowState() *YellowState {
	return &YellowState{
		durationSeconds: 2,
	}
}

func (this *YellowState) Change(context *trafficlight.TrafficLight) {
	fmt.Println("YELLOW light - Wait")
	time.Sleep(time.Duration(this.durationSeconds)*time.Second)
	context.SetLightState(NewRedState())
}
