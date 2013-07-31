package cslib

type Stack LinkedList

func (s *Stack) Push(value int) {
	((*LinkedList)(s)).Append(value)
}

func (s *Stack) Pop() int {
	l := (*LinkedList)(s)
	value := l.Tail.Value
	l.Remove(l.Length() - 1)
	return value
}

func (s *Stack) Empty() bool {
	return ((*LinkedList)(s)).Length() < 1
}
