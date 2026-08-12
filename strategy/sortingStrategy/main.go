package main

import (
	sortingservice "sortingstrategy/sortingService"
)

func main() {
	var input []int

	for i := 5000; i > 0; i-- {
		input = append(input, i)
	}

	// fmt.Println(sortingservice.NewTimedSort(sortingservice.NewQuickSort()).Sort(input))
	sortingservice.NewTimedSort(sortingservice.NewQuickSort()).Sort(input)
	sortingservice.NewTimedSort(sortingservice.NewMergeSort()).Sort(input)
	sortingservice.NewTimedSort(sortingservice.NewBubbleSort()).Sort(input)
}
