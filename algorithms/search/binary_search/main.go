package main

func BinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		midPos := (start + end) / 2
		if target < arr[midPos] {
			end = midPos - 1
		} else if target > arr[midPos] {
			start = midPos + 1
		} else {
			return midPos
		}
	}
	return -1
}
