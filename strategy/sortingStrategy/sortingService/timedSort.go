package sortingservice

import (
	"fmt"
	"time"
)

type TimedSort struct {
	inner SortingStrategy
}

func NewTimedSort(inner SortingStrategy) TimedSort {
	return TimedSort{
		inner: inner,
	}
}

func (this TimedSort) Sort(array []int) []int {
	now := time.Now()
	res := this.inner.Sort(array)
	elapsed := time.Since(now)
	fmt.Println("Time taken: ", elapsed)
	return res
}
