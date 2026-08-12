package sortingservice

import "fmt"

type SortingStrategy interface {
	Sort(array []int) []int
}

type Sorter struct {
	strategy SortingStrategy
}

func NewSorting(strategy SortingStrategy) *Sorter {
	return &Sorter{
		strategy: strategy,
	}
}

func (this *Sorter) SetStrategy(strategy SortingStrategy) {
	this.strategy = strategy
}

func (this *Sorter) Sort(array []int) {
	result := this.strategy.Sort(array)
	fmt.Println(result)
}
