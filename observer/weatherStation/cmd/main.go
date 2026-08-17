package main

import (
	"weatherstation/internals/model"
	"weatherstation/internals/service"
	"weatherstation/internals/service/observers"
)

func main() {
	w1 := model.WeatherData{
		Temperature: 32.4,
		Humidity: 61,
		Pressure: 682,
	}

	w2 := model.WeatherData{
		Temperature: 35.4,
		Humidity: 55,
		Pressure: 685,
	}

	o1 := observers.NewCurrentConditionDisplay()
	o2 := observers.NewStatisticsDisplay()

	weatherStation := service.NewWeatherStation()
	weatherStation.AddObserver(o1)
	weatherStation.AddObserver(&o2)

	weatherStation.SetMeasurements(w1)
	weatherStation.SetMeasurements(w2)

}
