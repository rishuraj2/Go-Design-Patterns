package trafficlight

type LightState interface {
	Change(context *TrafficLight)
}

type TrafficLight struct {
	state LightState
}

func NewTrafficLight(state LightState) *TrafficLight {
	return &TrafficLight{
		state: state,
	}
}

func (this *TrafficLight) SetLightState(state LightState) {
	this.state = state
}

func (this *TrafficLight) Change() {
	this.state.Change(this)
}
