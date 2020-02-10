package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	arr := []int64{30, 40, 52, 1, 9, 7, 10, -1}
	quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int64) {
	doQuickSort(arr, 0, int64(len(arr)-1))
}

func doQuickSort(arr []int64, first, last int64) {
	if first < last {
		pivot := pivot(arr, first, last)
		doQuickSort(arr, first, pivot-1)
		doQuickSort(arr, pivot+1, last)
	}
}

func pivot(arr []int64, first, last int64) int64 {
	pivot := randint64(first, last)
	swap(arr, pivot, last)
	for i := first; i < last; i++ {
		if arr[i] <= arr[last] {
			swap(arr, i, first)
			first++
		}
	}
	swap(arr, first, last)
	return first
}

func swap(arr []int64, i, j int64) {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
}

func randint64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}
