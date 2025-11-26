package main

import "slices"

type Queue[T any] struct {
	Elements []T
	CurIndex int
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{make([]T, 0, 1), 0}
}

func (q *Queue[T]) Enqueue(d T) {
	if q.CurIndex >= len(q.Elements) {
		q.Elements = slices.Grow(q.Elements, 1)
	}
	q.Elements = slices.Insert(q.Elements, len(q.Elements), d)
	q.CurIndex += 1
}

func (q *Queue[T]) Dequeue() T {
	dequeued := q.Elements[0]
	t := make([]T, 0, len(q.Elements)-1)
	t = q.Elements[1:]
	q.Elements = t
	q.CurIndex -= 1
	return dequeued
}

func (q *Queue[T]) GetElements() (r []T) {
	return q.Elements
}

// Generic type comparison too complicated rn.
// func (q *Queue[string]) Find(e string) int {
// 	var lower int = 0
// 	var upper int = len(q.Elements) - 1
// 	var mid int = lower + ((upper - lower) / 2)
// 	for lower <= upper {
// 		var val string = q.Elements[mid]
// 		if val == e {
// 			return mid
// 		} else if val < e {
// 			lower = upper
// 		} else if val > e {
// 		}
// 		mid = lower + ((upper - lower) / 2)
// 	}
// 	return -1
// }
