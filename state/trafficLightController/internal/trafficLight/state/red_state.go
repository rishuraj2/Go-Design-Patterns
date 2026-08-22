package state

import (
	"fmt"
	"time"
	trafficlight "trafficlightcontroller/internal/trafficLight"
)

type RedState struct {
	durationSeconds int
}

func NewRedState() *RedState {
	return &RedState{
		durationSeconds: 5,
	}
}

func (this *RedState) Change(context *trafficlight.TrafficLight) {
	fmt.Println("RED light - Stop")
	time.Sleep(time.Duration(this.durationSeconds)*time.Second)
	context.SetLightState(NewGreenState())
}
