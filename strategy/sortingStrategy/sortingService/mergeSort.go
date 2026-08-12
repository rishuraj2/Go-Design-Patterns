package sortingservice

import "fmt"

type MergeSort struct{}

func NewMergeSort() *MergeSort {
	return &MergeSort{}
}

func (this *MergeSort) Sort(array []int) []int {
	fmt.Println("using merge sort...")
	return this.mergeSort(array)
}

func (this *MergeSort) mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2

	left := this.mergeSort(array[:mid])
	right := this.mergeSort(array[mid:])

	return this.merge(left, right)
}

func (this *MergeSort) merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	if i < len(left) {
		result = append(result, left[i:]...)
	}
	
	if j < len(right) {
		result = append(right, right[j:]...)
	}

	return result
}
