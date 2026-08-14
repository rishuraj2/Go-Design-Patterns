package main

import (
	"fitnesstracker/internal/model"
	"fitnesstracker/internal/service"
	"fitnesstracker/internal/service/observer"
)

func main() {
	liveDispaly := observer.NewLiveActivityDisplay()
	goalNotifier := observer.NewGoalNotifier(1000, 60, 500)

	tracker := service.NewFitnessTracker()

	tracker.AddObserver(liveDispaly)
	tracker.AddObserver(goalNotifier)

	tracker.PushFitnessData(model.FitnessData{
		Steps: 900,
		ActiveMinutes: 50,
		Calories: 400,
	})

	tracker.PushFitnessData(model.FitnessData{
		Steps: 1000,
		ActiveMinutes: 55,
		Calories: 500,
	})
}
