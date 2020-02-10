package main

import "fmt"

func main() {
	result := make(chan []int)

	arr := []int{3000, 5, 120, -47, 281, 0, 3, 543, 31, 456, 786, 7, 43, 3, 221}

	go mergeSort(arr, result)

	merged := make([]int, len(arr))

	for i, r := range <-result {
		merged[i] = r
	}

	fmt.Println(merged)
}

func mergeSort(arr []int, result chan []int) {
	if len(arr) < 2 {
		result <- arr
		return
	}

	leftChan := make(chan []int)
	rightChan := make(chan []int)

	mid := len(arr) / 2

	go mergeSort(arr[:mid], leftChan)
	go mergeSort(arr[mid:], rightChan)

	left := <-leftChan
	right := <-rightChan

	result <- merge(left, right)
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
