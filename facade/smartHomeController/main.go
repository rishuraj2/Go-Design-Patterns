package main

import (
	"smarthomecontroller/controller"
	smarthomefacade "smarthomecontroller/smartHomeFacade"
)

func main() {
	light := controller.NewLight()
	thermostat := controller.NewThermostat()
	security := controller.NewSecurity()

	smartHome := smarthomefacade.NewSmartHomeFacade(&light, &security, &thermostat)

	smartHome.LeaveHome()

}
