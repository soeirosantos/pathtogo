package main

import (
	"errors"
)

// Kahn’s algorithm for Topological Sorting
func TopologicalSort(g *Graph) ([]int, error) {
	inDegree := make([]int, g.numOfVertices)

	for _, l := range g.graph {
		for _, j := range l {
			inDegree[j]++
		}
	}

	q := newQueue()

	for i := 0; i < g.numOfVertices; i++ {
		if inDegree[i] == 0 {
			q.push(i)
		}
	}

	count := 0

	topOrder := make([]int, g.numOfVertices)

	for q.size() > 0 {
		u, _ := q.poll()

		topOrder[count] = u

		for _, i := range g.graph[u] {
			inDegree[i]--
			if inDegree[i] == 0 {
				q.push(i)
			}
		}

		count++
	}

	if count != g.numOfVertices {
		return nil, errors.New("There exists a cycle in the graph")
	}

	return topOrder, nil
}

type Graph struct {
	graph         map[int][]int
	numOfVertices int
}

func NewGraph(numOfVertices int) *Graph {
	return &Graph{graph: make(map[int][]int), numOfVertices: numOfVertices}
}

func (g *Graph) AddEdge(u int, v int) {
	g.graph[u] = append(g.graph[u], v)
}

type queue struct {
	queue []int
	pos   int
}

func newQueue() *queue {
	return &queue{queue: make([]int, 10), pos: 0}
}

func (q *queue) push(v int) {
	q.queue[q.pos] = v
	q.pos++
}

func (q *queue) poll() (int, error) {
	if q.pos == 0 {
		return 0, errors.New("cannot poll empty queue")
	}
	var res int
	res, q.queue = q.queue[0], q.queue[1:]
	q.pos--
	return res, nil
}

func (q *queue) size() int {
	return q.pos
}
