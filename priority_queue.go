package iterutil

import (
	"cmp"
	"container/heap"
)

type PriorityQueue[T cmp.Ordered] struct {
	data genericsHeap[T]
}

type genericsHeap[T cmp.Ordered] []T

func (h genericsHeap[T]) Len() int {
	return len(h)
}

func (h genericsHeap[T]) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h genericsHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *genericsHeap[T]) Push(x interface{}) {
	*h = append(*h, x.(T))
}

func (h *genericsHeap[T]) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	pq := PriorityQueue[T]{}
	heap.Init(&pq.data)
	return &pq
}

func (pq *PriorityQueue[T]) Push(v T) {
	heap.Push(&pq.data, v)
}

func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(&pq.data).(T)
}

func (pq *PriorityQueue[T]) Peek() T {
	return pq.data[len(pq.data)-1]
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.data.Len()
}

func (pq *PriorityQueue[T]) Iter(yield func(T) bool) {
	for pq.Len() > 0 {
		if !yield(pq.Pop()) {
			break
		}
	}
}
