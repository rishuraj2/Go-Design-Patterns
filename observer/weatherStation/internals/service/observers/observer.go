package observers

import "weatherstation/internals/model"

type Observer interface {
	Notify(data model.WeatherData)
}
