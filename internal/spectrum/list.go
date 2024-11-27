package spectrum

type Node[T any] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}

type List[T any] struct {
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
}

func (list *List[T]) Concat(other *List[T]) {
	current := other.Head
	for current != nil {
		list.Append(current.Data)
		current = current.Next
	}
}

func (list *List[T]) ToArray() []T {
	arr := make([]T, 0)
	current := list.Head
	for current != nil {
		arr = append(arr, current.Data)
		current = current.Next
	}
	return arr
}
