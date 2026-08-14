package observer

import (
	"fitnesstracker/internal/model"
	"fmt"
)

type GoalNotifier struct {
	steps        int
	activeMinute int
	calories     int
}

func NewGoalNotifier(steps, activeMinutes, calories int) GoalNotifier {
	return GoalNotifier{
		steps: steps,
		activeMinute: activeMinutes,
		calories: calories,
	}
}

func (this GoalNotifier) Update(data model.FitnessData) {
	if this.steps == data.Steps {
		fmt.Printf("[Goal Achieved] Total steps of %d achieved\n", this.steps)
	}

	if this.activeMinute == data.ActiveMinutes {
		fmt.Printf("[Goal Achieved] Total active minutes of %d achieved\n", this.activeMinute)
	}

	if this.calories == data.Calories {
		fmt.Printf("[Goal Achieved] Total calories of %d achieved\n", this.calories)
	}
}
