package cslib

import "math"

type ArrayList struct {
	Data   []int
	Growth float64 // Rate at which slice grows when we overwrite its bounds.
}

func (list *ArrayList) Append(value int) {
	// We are going to implement our own append method to study what
	// different growth factors do to performance.

	// If we weren't doing our own list growth ratio, we would just do:
	// list.Data = append(list.Data, value)

	// Our own implementation of append:
	if len(list.Data) >= cap(list.Data) {
		oldData := list.Data
		newSize := int(math.Ceil(float64(len(list.Data)+1) * list.Growth))
		list.Data = make([]int, len(list.Data), newSize)
		copy(list.Data, oldData)
	}

	list.Data = list.Data[:len(list.Data)+1]
	list.Data[len(list.Data)-1] = value
}

func (list *ArrayList) Get(index int) int {
	return list.Data[index]
}

func (list *ArrayList) Remove(index int) {
	copy(list.Data[index:], list.Data[index+1:])
	list.Data = list.Data[:len(list.Data)-1]
}

func (list *ArrayList) Length() int {
	return len(list.Data)
}

func (list *ArrayList) Traverse(f func(int)) {
	for i := 0; i < len(list.Data); i++ {
		f(list.Data[i])
	}
}
