package main

import (
	"testing"
)

func TestTopologicalSort(t *testing.T) {

	g := NewGraph(6)

	g.AddEdge(5, 2)
	g.AddEdge(5, 0)
	g.AddEdge(4, 0)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)

	sorted, err := TopologicalSort(g)

	if err != nil {
		t.Error("Sort failed")
	}

	expected := []int{4, 5, 2, 0, 3, 1}

	for i, v := range expected {
		if sorted[i] != v {
			t.Error("Sort is wrong")
		}
	}
}

func TestPush(t *testing.T) {
	q := newQueue()

	q.push(2)
	q.push(6)
	q.push(-1)

	if q.size() != 3 {
		t.Errorf("Queue has wrong size, got %d want %d", q.size(), 3)
	}
}

func TestPoll(t *testing.T) {
	q := newQueue()

	tables := []struct {
		expected int
	}{
		{2},
		{6},
		{-1},
	}

	for _, table := range tables {
		q.push(table.expected)
	}

	if q.size() != 3 {
		t.Errorf("Queue has wrong size, got %d want %d", q.size(), 3)
	}

	for _, table := range tables {
		res, _ := q.poll()

		if res != table.expected {
			t.Errorf("Wrong result for poll, got %d want %d", res, table.expected)
		}
	}

	_, err := q.poll()

	if err == nil {
		t.Errorf("Error excpeted when polling am empty queue")
	}
}
