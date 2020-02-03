package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{5, 9, 23, 245, 342, 900, 1003, 2000}
	tables := []struct {
		target   int
		expected int
	}{
		{5, 0},
		{2000, 7},
		{342, 4},
		{9, 1},
		{1000, -1},
	}

	for _, table := range tables {
		actual := BinarySearch(nums, table.target)
		if actual != table.expected {
			t.Errorf("Found is wrong, got: %d, want: %d.", actual,
				table.expected)
		}

	}
}
