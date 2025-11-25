package main

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue[uint8]()
	q.Enqueue(128)
	if len(q.Elements) != 1 {
		t.Errorf("len(NewQueue) != %d", 2)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[uint8]()
	q.Enqueue(128)
	deq := q.Dequeue()
	if q.CurIndex != 0 {
		t.Errorf("len(Elements) = %d but CurIndex = %d", len(q.Elements), q.CurIndex)
	} else if deq != 128 {
		t.Errorf("Dequeue returned wrong value, expected %d got %d", 128, deq)
	}
}
