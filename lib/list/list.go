package list

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	Value T
	Next  *List[T]
}

func New[T any](value T) *List[T] {
	root := new(List[T])
	root.Value = value
	root.Next = nil
	return root
}

func FromSlice[T any](slice []T) *List[T] {
	var root *List[T]
	var last *List[T]

	for _, value := range slice {
		item := new(List[T])
		item.Value = value
		item.Next = nil

		if last != nil {
			last.Next = item
		}
		last = item

		if root == nil {
			root = item
		}
	}

	return root
}

func (l *List[T]) Len() int {
	node := l
	count := 0
	for node != nil {
		count++
		node = node.Next
	}
	return count
}

func (l *List[T]) Each(fn func(*List[T], T) bool) {
	node := l
	for node != nil {
		if !fn(node, node.Value) {
			return
		}
		node = node.Next
	}
}

func (l *List[T]) String() string {
	type Stringer interface {
		String() string
	}

	var buf strings.Builder
	buf.WriteString("[")
	for n := l; n != nil; n = n.Next {
		if buf.Len() > 1 {
			buf.WriteString(", ")
		}
		if str, ok := any(n.Value).(Stringer); ok {
			buf.WriteString(str.String())
		} else {
			fmt.Fprintf(&buf, "%v", n.Value)
		}
	}
	buf.WriteString("]")
	return buf.String()
}
