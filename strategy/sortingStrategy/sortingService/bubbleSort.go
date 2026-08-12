package sortingservice

import "fmt"

type BubbleSort struct{}

func NewBubbleSort() *BubbleSort {
	return &BubbleSort{}
}

func (this *BubbleSort) Sort(array []int) []int {
	fmt.Println("using bubble sort...")

	arrLength := len(array)

	for i := 0; i < arrLength-1; i++ {
		swapped := false

		for j := 0; j < arrLength-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return array
}
