package stack

func New[T any]() Stack[T] {
	return Stack[T]{}
}

type Stack[T any] []T

func (s Stack[T]) Slice() []T {
	c := make([]T, len(s))
	copy(c, s)
	return c
}

func (s Stack[T]) Len() int {
	return len(s)
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() (r T) {
	if len(*s) != 0 {
		i := len(*s) - 1
		r, *s = (*s)[i], (*s)[:i]
	}
	return
}

func (s Stack[T]) Top() T {
	i := len(s) - 1
	if i >= 0 {
		return s[i]
	}
	var zero T
	return zero
}
