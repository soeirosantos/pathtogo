package main

import (
	"fmt"
)

func main() {
	fmt.Println(mergeSort([]int{20, 30, 1, 200, 3, -6, 90000}))
}

func mergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {

	merged := []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	merged = append(merged, left...)
	merged = append(merged, right...)

	return merged
}
