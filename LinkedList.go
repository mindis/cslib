package cslib

type LinkedListNode struct {
	Previous *LinkedListNode
	Next     *LinkedListNode
	Value    int
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
	Len  int
}

func (list *LinkedList) Append(value int) {
	node := &LinkedListNode{Value: value}

	if list.Head == nil {
		list.Head = node
	}

	if list.Tail == nil {
		node.Previous = nil
		list.Tail = node
	} else {
		node.Previous = list.Tail
		list.Tail.Next = node
		list.Tail = node
	}

	list.Len += 1
}

func (list *LinkedList) Get(index int) int {
	node := list.Head
	for index > 0 {
		node = node.Next
		index--
	}
	return node.Value
}

func (list *LinkedList) Remove(index int) {
	node := list.Head
	for index > 0 {
		node = node.Next
		index--
	}

	if list.Head == node {
		list.Head = node.Next
	}

	if list.Tail == node {
		list.Tail = node.Previous
	}

	if node.Previous != nil {
		node.Previous.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Previous = node.Previous
	}

	node.Previous = nil
	node.Next = nil
	list.Len -= 1
}

func (list *LinkedList) Length() int {
	return list.Len
}

func (list *LinkedList) Traverse(f func(int)) {
	for node := list.Head; node != nil; node = node.Next {
		f(node.Value)
	}
}

func (list *LinkedList) Search(value int) *LinkedListNode {
	for node := list.Head; node != nil; node = node.Next {
		// We assume the list is sorted.
		if node.Value > value {
			break
		} else if node.Value == value {
			return node
		}
	}

	return nil
}
