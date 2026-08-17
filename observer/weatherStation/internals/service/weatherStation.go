package service

import (
	"fmt"
	"weatherstation/internals/model"
	"weatherstation/internals/service/observers"
)

type WeatherStation struct {
	weatherData model.WeatherData
	observers   []observers.Observer
}

func NewWeatherStation() WeatherStation {
	return WeatherStation{}
}

func (this *WeatherStation) AddObserver(observer observers.Observer) {
	this.observers = append(this.observers, observer)
}

func (this *WeatherStation) RemoveObserver(observer observers.Observer) {
	for i, o := range this.observers {
		if o == observer {
			this.observers = append(this.observers[:i], this.observers[i+1:]...)
		}
	}
}

func (this *WeatherStation) SetMeasurements(data model.WeatherData) {
	this.weatherData = data
	fmt.Printf("[Weather Station] New readings recieved. Temperature: %f, Humidity: %f, Pressure: %f\n", data.Temperature, data.Humidity, data.Pressure)
	this.notifyObserver()
}

func (this WeatherStation) notifyObserver() {
	for _, observer := range this.observers {
		observer.Notify(this.weatherData)
	}
}
