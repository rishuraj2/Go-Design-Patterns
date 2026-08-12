package sortingservice

import "fmt"

type QuickSort struct{}

func NewQuickSort() *QuickSort {
	return &QuickSort{}
}

func (this *QuickSort) Sort(array []int) []int {
	fmt.Println("using quick sort...")

	if len(array) <= 1 {
		return array
	}
	
	this.quickSortRange(array, 0, len(array)-1)

	return array
}

func (this *QuickSort) quickSortRange(array []int, low, high int) {
	if low < high {
		pIdx := this.partition(array, low, high)

		this.quickSortRange(array, low, pIdx-1)
		this.quickSortRange(array, pIdx+1, high)
	}
}

func (this *QuickSort) partition(array []int, low, high int) int {
	pivot := array[high]
	
	i := low - 1

	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[high] = array[high], array[i+1]
	
	return i + 1
}

