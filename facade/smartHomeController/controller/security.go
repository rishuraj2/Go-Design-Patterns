package controller

type Security struct {
	state bool
}

func NewSecurity() Security {
	return Security{
		state: false,
	}
}

func (this *Security) SetState(state bool) {
	this.state = state
}

func (this Security) GetState() string {
	if this.state {
		return "armed"
	}

	return "unarmed"
}
