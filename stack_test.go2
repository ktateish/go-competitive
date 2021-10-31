package main

// vim:set ft=go:

import "testing"

func TestStack(t *testing.T) {
	s := NewStack[int]()
	inputs := []int{1, 2, 3, 4, 5}
	wants := []int{5, 4, 3, 2, 1}

	for i, v := range inputs {
		if want, got := i, s.Len(); want != got {
			t.Fatalf("s.Len() returns wrong value: want=%d got=%d\n", want, got)
		}
		s.Push(v)
	}
	if want, got := len(inputs), s.Len(); want != got {
		t.Fatalf("s.Len() returns wrong value: want=%d got=%d\n", want, got)
	}

	for i := 0; 0 < s.Len(); i++ {
		got := s.Pop()
		want := wants[i]
		if got != want {
			t.Fatalf("s.Pop() returns wrong value: want=%d got=%d\n", want, got)
		}
	}
}
