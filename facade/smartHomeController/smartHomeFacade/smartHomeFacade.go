package smarthomefacade

import (
	"fmt"
	"smarthomecontroller/controller"
)

type SmartHomeFacade struct {
	light      *controller.Light
	security   *controller.Security
	thermostat *controller.Thermostat
}

func NewSmartHomeFacade(light *controller.Light, security *controller.Security, thermostat *controller.Thermostat) SmartHomeFacade {
	return SmartHomeFacade{
		light:      light,
		security:   security,
		thermostat: thermostat,
	}
}

func (this *SmartHomeFacade) LeaveHome() {
	this.light.SetState(false)
	this.security.SetState(true)
	this.thermostat.SetState(false)

	fmt.Printf("[Leaving Home] Light: %s, Thermostat: %s, Security: %s\n", this.light.GetStateString(), this.thermostat.GetStateString(), this.security.GetState())
}

func (this *SmartHomeFacade) ArriveHome() {
	this.light.SetState(true)
	this.security.SetState(false)
	this.thermostat.SetState(true)

	fmt.Printf("[Arriving Home] Light: %s, Thermostat: %s, Security: %s\n", this.light.GetStateString(), this.thermostat.GetStateString(), this.security.GetState())
}
