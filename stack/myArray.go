package stack

import "errors"

type Array[C any] struct {
	length int64
	list   []C
}

func (array *Array[C]) Len() int64 {
	return array.length
}

func (array Array[C]) Get(index int64) any {
	return array.list[index]
}
func (array *Array[C]) Push(item C) {
	array.list = append(array.list, item)
	array.length++
}

func (array *Array[C]) Pop() (any, error) {
	item := array.list[array.length-1]
	if array.length <= 0 {
		return nil, errors.New("array is empty")
	}
	array.list = array.list[:array.length-1]
	array.length--
	return item, nil
}

func (array *Array[C]) Delete(index int64) (any, error) {
	item := array.list[index]
	error := array.shiftItem(index)
	return item, error
}

func (array *Array[C]) shiftItem(index int64) error {
	if array.length <= 0 || index >= array.length {
		return errors.New("array is empty")
	}
	for i := index; i < array.length-1; i++ {
		array.list[i] = array.list[i+1]
	}
	array.list = array.list[:array.length-1]
	array.length--
	return nil
}

func New() *Array[any] {
	return &Array[any]{
		length: 0,
		list:   []any{},
	}
}
