package main

import (
	"fmt"
	"temperatureconverter/adapter"
	legacyconverter "temperatureconverter/legacyConverter"
)

func main() {
	fah := legacyconverter.NewFahrenheitSensor(96)
	cel := adapter.NewCelsiusSensor(fah)
	fmt.Printf("%.2f\n", cel.GetTemperature())
}
