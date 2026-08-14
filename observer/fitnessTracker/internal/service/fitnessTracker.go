package service

import (
	"fitnesstracker/internal/model"
	"fmt"
)

type FitnessDataObserver interface {
	Update(data model.FitnessData)
}

type FitnessTracker struct {
	data      model.FitnessData
	observers []FitnessDataObserver
}

func NewFitnessTracker() *FitnessTracker {
	return &FitnessTracker{}
}

func (this *FitnessTracker) AddObserver(observer FitnessDataObserver) {
	this.observers = append(this.observers, observer)
}

func (this *FitnessTracker) RemoveObserver(observer FitnessDataObserver) {
	for i, o := range this.observers {
		if o == observer {
			this.observers = append(this.observers[:i], this.observers[i+1:]...)
		}
	}
}

func (this *FitnessTracker) notifyObservers() {
	for _, observer := range this.observers {
		observer.Update(this.data)
	}
}

func (this *FitnessTracker) PushFitnessData(data model.FitnessData) {
	this.data = data
	fmt.Printf("[Fitness Tracker] New data recieved | Steps: %d, Active Minutes: %d, Calories: %d\n", this.data.Steps, this.data.ActiveMinutes, this.data.Calories)
	this.notifyObservers()
}

func (this *FitnessTracker) Reset() {
	this.data = model.FitnessData{}
	fmt.Println("[Fitness Tracker] Data reset performed!")
}
