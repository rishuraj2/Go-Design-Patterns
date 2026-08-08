package adapter

import legacyconverter "temperatureconverter/legacyConverter"

type CelsiusSensor struct {
	fahrenheitSensor legacyconverter.FahrenheitSensor
}

func NewCelsiusSensor(fahren legacyconverter.FahrenheitSensor) CelsiusSensor {
	return CelsiusSensor{
		fahrenheitSensor: fahren,
	}
}

func (c CelsiusSensor) GetTemperature() float64 {
	fah := c.fahrenheitSensor.ReadFahrenheit()
	return (fah - 32) * (5.0/9.0)
}
