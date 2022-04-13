package model

type Stack[T any] struct {
	entries []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		entries: nil,
	}
}

func (s *Stack[T]) Push(val T) {
	s.entries = append(s.entries, val)
}

func (s *Stack[T]) Pop() (t T) {
	size := s.Size()
	if size == 0 {
		return t
	}

	t = s.entries[size-1]
	s.entries = s.entries[:size-1]
	return t
}

func (s *Stack[T]) Peek() (t T) {
	size := s.Size()
	if size == 0 {
		return t
	}
	return s.entries[size-1]
}

func (s *Stack[T]) Size() int {
	return len(s.entries)
}
