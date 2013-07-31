package cslib

type Queue LinkedList

func (q *Queue) Enqueue(value int) {
	((*LinkedList)(q)).Append(value)
}

func (q *Queue) Dequeue() int {
	l := (*LinkedList)(q)
	value := l.Head.Value
	l.Remove(0)
	return value
}

func (q *Queue) Empty() bool {
	return ((*LinkedList)(q)).Length() < 1
}
