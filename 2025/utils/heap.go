package utils

import (
	"container/heap"
)

type element[T any] struct {
	value    *T
	priority int
	index    int
}

type store[T any] []element[T]

func (s store[T]) Len() int {
	return len(s)
}

func (s store[T]) Less(i, j int) bool {
	return s[i].priority < s[j].priority
}

func (s store[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	s[i].index = i
	s[j].index = j
}

func (s *store[T]) Push(x any) {
	e := x.(element[T])
	e.index = len(*s)

	*s = append(*s, e)
}

func (s *store[T]) Pop() any {
	e := (*s)[len(*s)-1]
	e.index = -1

	*s = (*s)[0 : len(*s)-1]
	return e
}

type Heap[T any] struct {
	s *store[T]
}

func NewHeap[T any]() Heap[T] {
	h := Heap[T]{}
	h.s = &store[T]{}

	heap.Init(h.s)
	return h
}

func NewHeapFrom[T any](items map[*T]int) Heap[T] {
	h := Heap[T]{}
	h.s = &store[T]{}

	heap.Init(h.s)
	for value, priority := range items {
		heap.Push(h.s, element[T]{value: value, priority: priority})
	}

	return h
}

func (h Heap[T]) Len() int {
	return h.s.Len()
}

func (h Heap[T]) Empty() bool {
	return h.Len() == 0
}

func (h Heap[T]) Push(value *T, priority int) {
	heap.Push(h.s, element[T]{value: value, priority: priority})
}

func (h Heap[T]) Pop() *T {
	return heap.Pop(h.s).(element[T]).value
}
