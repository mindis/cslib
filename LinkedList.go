package cslib

type Node struct {
	Previous *Node
	Next     *Node
	Value    int
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func (list *LinkedList) Append(value int) {
	node := &Node{Value: value}

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
	node := list.Head
	for node != nil {
		f(node.Value)
		node = node.Next
	}
}

func (list *LinkedList) Search(value int) int {
	i := 0
	node := list.Head

	for node != nil {
		// We assume the list is sorted.
		if node.Value > value {
			break
		} else if node.Value == value {
			return i
		}

		node = node.Next
		i++
	}

	return -1
}
