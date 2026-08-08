package legacyconverter

type FahrenheitSensor struct {
	temperature float64
}

func NewFahrenheitSensor(fahren float64) FahrenheitSensor {
	return FahrenheitSensor{
		temperature: fahren,
	}
}

func (f FahrenheitSensor) ReadFahrenheit() float64 {
	return f.temperature
}
