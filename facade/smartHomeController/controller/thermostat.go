package controller

type Thermostat struct {
	state bool
}

func NewThermostat() Thermostat {
	return Thermostat{
		state: false,
	}
}

func (this *Thermostat) SetState(state bool) {
	this.state = state
}

func (this Thermostat) GetStateString() string {
	if this.state {
		return "heating"
	}

	return "off"
}
