package main

import (
	"fmt"
)

func main() {

	result := make(chan []int64)
	arr := []int64{3000, 5, 120, -47, 281, 0, 3, 543, 31, 456, 786, 7, 43, 3, 221}
	go mergeSort(arr, result)

	sorted := make([]int64, len(arr))
	for i, r := range <-result {
		sorted[i] = r
	}

	fmt.Println(sorted)
}

func mergeSort(arr []int64, result chan []int64) {
	if len(arr) < 2 {
		result <- arr
		return
	}

	leftChan := make(chan []int64)
	rightChan := make(chan []int64)

	mid := len(arr) / 2

	go mergeSort(arr[:mid], leftChan)
	go mergeSort(arr[mid:], rightChan)

	left := <-leftChan
	right := <-rightChan

	result <- merge(left, right)
}

func merge(left, right []int64) []int64 {

	merged := []int64{}

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
