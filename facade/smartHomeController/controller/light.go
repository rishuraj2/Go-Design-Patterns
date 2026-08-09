package controller

type Light struct {
	state bool
}

func NewLight() Light {
	return Light{
		state: false,
	}
}

func (this *Light) SetState(state bool) {
	this.state = state
}

func (this Light) GetStateString() string {
	if this.state {
		return "on"
	}

	return "off"
}
