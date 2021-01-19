package main

import (
	"errors"
	"fmt"
)

// This is an instructional example of a Queue backed by two Stacks

func main() {

	q := newQueue(10)

	q.push(3)
	q.push(5)
	q.push(1)

	fmt.Println(q)

	fmt.Println(q.poll())
	fmt.Println(q.poll())
	fmt.Println(q.poll())
	
	fmt.Println(q)
	
	q.push(1)
	fmt.Println(q.poll())

	q.push(2)
	q.push(3)
	fmt.Println(q.poll())
	fmt.Println(q.poll())
}

type queue struct {
	in  stack
	out stack
}

func newQueue(capacity int) queue {
	return queue{newStack(capacity), newStack(capacity)}
}

func (q *queue) push(el int) {
	q.in.push(el)
}

func (q *queue) poll() (int, error) {
	if q.out.size() == 0 {
		for q.in.size() > 0 {
			el, _ := q.in.pop()
			q.out.push(el)
		}
	}
	p, err := q.out.pop()
	if err != nil {
		return p, errors.New("cannot poll from empty queue.")
	}
	return p, nil
}

type stack struct {
	arr []int
	pos int
}

func newStack(capacity int) stack {
	return stack{arr: make([]int, capacity)}
}

func (s *stack) push(el int) {
	if len(s.arr) == s.pos {
		new := make([]int, 2*len(s.arr))
		copy(new, s.arr)
		s.arr = new
	}
	s.arr[s.pos] = el
	s.pos++
}

func (s *stack) pop() (int, error) {
	if s.pos == 0 {
		return -1, errors.New("cannot pop from empty stack.")
	}
	s.pos--
	el := s.arr[s.pos]
	s.arr[s.pos] = 0
	return el, nil
}

func (s stack) size() int {
	return s.pos
}
