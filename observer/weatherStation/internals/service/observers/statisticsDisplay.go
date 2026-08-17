package observers

import (
	"fmt"
	"weatherstation/internals/model"
)

type StatisticsDisplay struct{
	temperatures []float64
}

func NewStatisticsDisplay() StatisticsDisplay {
	return StatisticsDisplay{}
}

func (this *StatisticsDisplay) Notify(data model.WeatherData) {
	this.storeTemperature(data.Temperature)
	fmt.Printf("[Statistics Display] Average Temperature: %f\n", this.calcAverage())
}

func (this *StatisticsDisplay) storeTemperature(temperature float64) {
	this.temperatures = append(this.temperatures, temperature)
}

func (this StatisticsDisplay) calcAverage() float64 {
	totalTemp := 0.0
	for _, temp := range this.temperatures {
		totalTemp += temp
	}
	return totalTemp / float64(len(this.temperatures))
}
