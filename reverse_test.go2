package main

// vim:set ft=go:

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title string
		input []int
		want  []int
	}{
		{
			title: "simple usage",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			title: "zero length",
			input: []int{},
			want:  []int{},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run("Reverse:"+ts.title, func(t *testing.T) {
			Reverse(ts.input)
			if !reflect.DeepEqual(ts.want, ts.input) {
				t.Errorf("Reverse() failed to reverse : want=%v got=%v\n", ts.want, ts.input)
			}
		})
	}
}

func TestAReverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title string
		input []int
		want  []int
	}{
		{
			title: "simple usage",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			title: "zero length",
			input: []int{},
			want:  []int{},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run("Reverse:"+ts.title, func(t *testing.T) {
			got := AReverse(ts.input)
			if !reflect.DeepEqual(ts.want, got) {
				t.Errorf("AReverse() failed to reverse : want=%v got=%v\n", ts.want, got)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title  string
		input0 []int
		input1 int
		want   []int
	}{
		{
			title:  "simple usage",
			input0: []int{1, 2, 3, 4, 5},
			input1: 2,
			want:   []int{3, 4, 5, 1, 2},
		},
		{
			title:  "0 rotate",
			input0: []int{1, 2, 3, 4, 5},
			input1: 0,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			title:  "over length",
			input0: []int{1, 2, 3, 4, 5},
			input1: 13,
			want:   []int{4, 5, 1, 2, 3},
		},
		{
			title:  "zero length slice",
			input0: []int{},
			input1: 20,
			want:   []int{},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run("Rotate:"+ts.title, func(t *testing.T) {
			Rotate(ts.input0, ts.input1)
			if !reflect.DeepEqual(ts.want, ts.input0) {
				t.Errorf("Rotate() failed to reverse : want=%v got=%v\n", ts.want, ts.input0)
			}
		})
	}
}

func TestRotateLeft(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title  string
		input0 []int
		input1 int
		want   []int
	}{
		{
			title:  "simple usage",
			input0: []int{1, 2, 3, 4, 5},
			input1: 2,
			want:   []int{3, 4, 5, 1, 2},
		},
		{
			title:  "0 rotate",
			input0: []int{1, 2, 3, 4, 5},
			input1: 0,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			title:  "over length",
			input0: []int{1, 2, 3, 4, 5},
			input1: 13,
			want:   []int{4, 5, 1, 2, 3},
		},
		{
			title:  "zero length slice",
			input0: []int{},
			input1: 20,
			want:   []int{},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run("RotateLeft:"+ts.title, func(t *testing.T) {
			RotateLeft(ts.input0, ts.input1)
			if !reflect.DeepEqual(ts.want, ts.input0) {
				t.Errorf("RotateLeft() failed to reverse : want=%v got=%v\n", ts.want, ts.input0)
			}
		})
	}
}

func TestRotateRight(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title  string
		input0 []int
		input1 int
		want   []int
	}{
		{
			title:  "simple usage",
			input0: []int{1, 2, 3, 4, 5},
			input1: 2,
			want:   []int{4, 5, 1, 2, 3},
		},
		{
			title:  "0 rotate",
			input0: []int{1, 2, 3, 4, 5},
			input1: 0,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			title:  "over length",
			input0: []int{1, 2, 3, 4, 5},
			input1: 13,
			want:   []int{3, 4, 5, 1, 2},
		},
		{
			title:  "zero length slice",
			input0: []int{},
			input1: 20,
			want:   []int{},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run("Rotate:"+ts.title, func(t *testing.T) {
			RotateRight(ts.input0, ts.input1)
			if !reflect.DeepEqual(ts.want, ts.input0) {
				t.Errorf("RotateRight() failed to reverse : want=%v got=%v\n", ts.want, ts.input0)
			}
		})
	}
}

func TestARotate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title  string
		input0 []int
		input1 int
		want   []int
	}{
		{
			title:  "simple usage",
			input0: []int{1, 2, 3, 4, 5},
			input1: 2,
			want:   []int{3, 4, 5, 1, 2},
		},
		{
			title:  "0 rotate",
			input0: []int{1, 2, 3, 4, 5},
			input1: 0,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			title:  "over length",
			input0: []int{1, 2, 3, 4, 5},
			input1: 13,
			want:   []int{4, 5, 1, 2, 3},
		},
		{
			title:  "zero length slice",
			input0: []int{},
			input1: 20,
			want:   []int{},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run("ARotate:"+ts.title, func(t *testing.T) {
			got := ARotate(ts.input0, ts.input1)
			if !reflect.DeepEqual(ts.want, got) {
				t.Errorf("ARotate() failed to reverse : want=%v got=%v\n", ts.want, got)
			}
		})
	}
}

func TestARotateLeft(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title  string
		input0 []int
		input1 int
		want   []int
	}{
		{
			title:  "simple usage",
			input0: []int{1, 2, 3, 4, 5},
			input1: 2,
			want:   []int{3, 4, 5, 1, 2},
		},
		{
			title:  "0 rotate",
			input0: []int{1, 2, 3, 4, 5},
			input1: 0,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			title:  "over length",
			input0: []int{1, 2, 3, 4, 5},
			input1: 13,
			want:   []int{4, 5, 1, 2, 3},
		},
		{
			title:  "zero length slice",
			input0: []int{},
			input1: 20,
			want:   []int{},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run("ARotateLeft:"+ts.title, func(t *testing.T) {
			got := ARotateLeft(ts.input0, ts.input1)
			if !reflect.DeepEqual(ts.want, got) {
				t.Errorf("ARotateLeft() failed to reverse : want=%v got=%v\n", ts.want, got)
			}
		})
	}
}

func TestARotateRight(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title  string
		input0 []int
		input1 int
		want   []int
	}{
		{
			title:  "simple usage",
			input0: []int{1, 2, 3, 4, 5},
			input1: 2,
			want:   []int{4, 5, 1, 2, 3},
		},
		{
			title:  "0 rotate",
			input0: []int{1, 2, 3, 4, 5},
			input1: 0,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			title:  "over length",
			input0: []int{1, 2, 3, 4, 5},
			input1: 13,
			want:   []int{3, 4, 5, 1, 2},
		},
		{
			title:  "zero length slice",
			input0: []int{},
			input1: 20,
			want:   []int{},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run("ARotate:"+ts.title, func(t *testing.T) {
			got := ARotateRight(ts.input0, ts.input1)
			if !reflect.DeepEqual(ts.want, got) {
				t.Errorf("ARotateRight() failed to reverse : want=%v got=%v\n", ts.want, got)
			}
		})
	}
}
