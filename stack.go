package iterutil

import "container/list"

type Stack[T any] struct {
	data list.List
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: list.List{},
	}
}

func (s *Stack[T]) Push(v T) {
	s.data.PushBack(v)
}

func (s *Stack[T]) Pop() T {
	e := s.data.Back()
	if e == nil {
		var zero T
		return zero
	}

	s.data.Remove(e)
	return e.Value.(T)
}

func (s *Stack[T]) Peek() T {
	e := s.data.Back()
	if e == nil {
		var zero T
		return zero
	}

	return e.Value.(T)
}

func (s *Stack[T]) Len() int {
	return s.data.Len()
}

func (s *Stack[T]) Iter(yield func(T) bool) {
	for e := s.data.Back(); e != nil; e = s.data.Back() {
		s.data.Remove(e)

		yield(e.Value.(T))
	}
}
