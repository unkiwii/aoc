package list

import (
	"strings"
)

type Stringer interface {
	String() string
}

type List[T Stringer] struct {
	Value T
	Next  *List[T]
}

func FromSlice[T Stringer](slice []T) *List[T] {
	var root *List[T]
	var last *List[T]

	for _, v := range slice {
		item := new(List[T])
		item.Value = v
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

func (l *List[T]) String() string {
	var buf strings.Builder
	for n := l; n != nil; n = n.Next {
		if buf.Len() != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(n.Value.String())
	}
	return buf.String()
}
