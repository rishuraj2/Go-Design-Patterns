package observers

import (
	"fmt"
	"weatherstation/internals/model"
)

type CurrentConditionDisplay struct{}

func NewCurrentConditionDisplay() CurrentConditionDisplay {
	return CurrentConditionDisplay{}
}

func (this CurrentConditionDisplay) Notify(data model.WeatherData) {
	fmt.Printf("[Current Condition Display] Temperature: %f, Humidity: %f, Pressure: %f\n", data.Temperature, data.Humidity, data.Pressure)
}
