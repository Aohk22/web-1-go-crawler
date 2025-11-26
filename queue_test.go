package main

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue[uint8]()
	q.Enqueue(128)
	if len(q.elements) != 1 {
		t.Errorf("len(NewQueue) != %d", 2)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[uint8]()
	q.Enqueue(128)
	deq, _ := q.Dequeue()
	if q.GetLength() != 0 {
	} else if deq != 128 {
		t.Errorf("Dequeue returned wrong value, expected %d got %d", 128, deq)
	}
}

func TestFind(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("hellow")
	got := q.Exists("hellow")
	if got != true {
		t.Errorf("Expected true, got %t", got)
	}
}
