package cslib

// This interface defines how a List type should behave.
// We use integer data here for purposes of simplicity.
type List interface {
	Append(value int)
	Get(index int) int
	Remove(index int)
	Length() int
	Traverse(func(int))
}
