package main

import (
	"reflect"
	"testing"
)

// vim:set ft=go:

func TestLIS(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title string
		input []int
		want  []int
	}{
		{
			title: "2 6 4 5 7 1 3",
			input: []int{2, 6, 4, 5, 7, 1, 3},
			want:  []int{1, 2, 2, 3, 4, 1, 2},
		},
		{
			title: "zero length slice",
			input: []int{},
			want:  []int{},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run("LIS:"+ts.title, func(t *testing.T) {
			got := LIS(ts.input, Less[int])
			if !reflect.DeepEqual(ts.want, got) {
				t.Errorf("LIS(%v) returns wrong results: want=%v got=%v\n", ts.input, ts.want, got)
			}
		})
	}

}

func naiveLIS[T OrderedNumeric](a []T) []int {
	res := make([]int, len(a))
	for i := range a {
		found := false
		for j := i - 1; 0 <= j; j-- {
			if a[j] < a[i] {
				res[i] = res[j] + 1
				found = true
			}
		}
		if !found {
			res[i] = 1
		}
	}
	return res
}

func TestLISRandom(t *testing.T) {
	N := 20
	a := make([]int, 20)
	for i := range a {
		a[i] = i + 1
	}
	for tt := 0; tt < 100; tt++ {
		Shuffle(a)
		want := naiveLIS(a)
		got := LIS(a)
		if !reflect.DeepEqual(ts.want, got) {
			t.Errorf("LIS(%v) returns wrong results: want=%v got=%v\n", a, ts.want, got)
		}
	}
}
