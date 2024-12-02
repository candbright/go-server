package core

type Node[T any] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}

type List[T any] struct {
	Len  int
	Head *Node[T]
	Tail *Node[T]
}

func NewList[T any](data ...T) *List[T] {
	list := &List[T]{}
	for _, datum := range data {
		list.Append(datum)
	}
	return list
}

func (list *List[T]) Append(data T) {
	newNode := &Node[T]{Data: data}

	if list.Tail == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		newNode.Prev = list.Tail
		list.Tail = newNode
	}
	list.Len++
}

func (list *List[T]) Concat(other *List[T]) {
	current := other.Head
	for current != nil {
		list.Append(current.Data)
		current = current.Next
	}
}

func (list *List[T]) ToArray() []T {
	arr := make([]T, list.Len)
	current := list.Head
	for i := 0; i < list.Len; i++ {
		arr[i] = current.Data
		current = current.Next
	}
	return arr
}

func (list *List[T]) ForRange(f func(data T)) {
	current := list.Head
	for current != nil {
		f(current.Data)
		current = current.Next
	}
}
