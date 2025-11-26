package main

import "errors"
import "slices"

type Queue[T comparable] struct {
	elements []T
}

/* Initialize slice with capacity of 1. */
func NewQueue[T comparable]() Queue[T] {
	return Queue[T]{make([]T, 0, 1)}
}

func (q *Queue[T]) Enqueue(element T) {
	if len(q.elements) == cap(q.elements) {
		q.elements = slices.Grow(q.elements, 1)
	}
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() (deq T, err error) {
	if len(q.elements) == 0 {
		return deq, errors.New("Cannot dequeue from empty queue.")
	}

	deq = q.elements[0]
	copy(q.elements, q.elements[1:])
	q.elements = q.elements[:len(q.elements)-1]
	return deq, nil
}

/* Return queue as array. */
func (q *Queue[T]) GetElements() []T {
	return q.elements
}

func (q *Queue[T]) GetLength() int {
	return len(q.elements)
}

func (q *Queue[T]) Exists(e T) bool {
	return slices.Contains(q.elements, e)
}
