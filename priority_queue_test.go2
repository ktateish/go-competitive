package main

// vim:set ft=go:

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	N := 100
	want := make([]int, N)
	Iota(want, 1)
	input := make([]int, N)
	copy(input, want)
	Shuffle(input)

	pq := NewPriorityQueue(Less[int])
	for i := range input {
		pq.Push(input[i])
		want := i + 1
		got := pq.Len()
		if want != got {
			t.Fatalf("PriorityQueue.Len() returns wrong value aflter Push(%d): want=%d got=%d\n", input[i], want, got)
		}
	}
	for i := 0; i < N && !pq.Empty(); i++ {
		got := pq.Pop()
		if want[i] != got {
			t.Fatalf("PriorityQueue returns wrong value for %d-th pop(): want=%d got=%d\n", i, want[i], got)
		}
	}
}
