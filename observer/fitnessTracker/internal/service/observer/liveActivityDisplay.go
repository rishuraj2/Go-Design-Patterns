package observer

import (
	"fitnesstracker/internal/model"
	"fmt"
)

type LiveActivityDispaly struct{}

func NewLiveActivityDisplay() LiveActivityDispaly {
	return LiveActivityDispaly{}
}

func (this LiveActivityDispaly) Update(data model.FitnessData) {
	fmt.Printf(`
	------------ Live Update ------------
	Steps| %d
	Active Minutes| %d
	Calories| %d
	`, data.Steps, data.ActiveMinutes, data.Calories)
	fmt.Println()
}
