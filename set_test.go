package main

// vim:set ft=go:

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSetasic(t *testing.T) {
	tests := []struct {
		op    string
		input int
		want  bool
	}{
		{"Insert", 5, true},
		{"FindValue", 4, false},
		{"FindValue", 5, true},
		{"FindValue", 6, false},
		{"Insert", 5, false},

		{"Insert", 2, true},
		{"FindValue", 1, false},
		{"FindValue", 2, true},
		{"FindValue", 3, false},
		{"FindValue", 4, false},
		{"FindValue", 5, true},
		{"FindValue", 6, false},
		{"Insert", 2, false},

		{"Insert", 7, true},
		{"FindValue", 1, false},
		{"FindValue", 2, true},
		{"FindValue", 3, false},
		{"FindValue", 4, false},
		{"FindValue", 5, true},
		{"FindValue", 6, false},
		{"FindValue", 7, true},
		{"FindValue", 8, false},
		{"Insert", 7, false},

		{"Insert", 4, true},
		{"FindValue", 1, false},
		{"FindValue", 2, true},
		{"FindValue", 3, false},
		{"FindValue", 4, true},
		{"FindValue", 5, true},
		{"FindValue", 6, false},
		{"FindValue", 7, true},
		{"FindValue", 8, false},
		{"Insert", 4, false},

		{"Insert", 6, true},
		{"FindValue", 1, false},
		{"FindValue", 2, true},
		{"FindValue", 3, false},
		{"FindValue", 4, true},
		{"FindValue", 5, true},
		{"FindValue", 6, true},
		{"FindValue", 7, true},
		{"FindValue", 8, false},
		{"Insert", 6, false},

		{"Insert", 3, true},
		{"FindValue", 1, false},
		{"FindValue", 2, true},
		{"FindValue", 3, true},
		{"FindValue", 4, true},
		{"FindValue", 5, true},
		{"FindValue", 6, true},
		{"FindValue", 7, true},
		{"FindValue", 8, false},
		{"Insert", 3, false},

		{"Remove", 5, true},
		{"FindValue", 1, false},
		{"FindValue", 2, true},
		{"FindValue", 3, true},
		{"FindValue", 4, true},
		{"FindValue", 5, false},
		{"FindValue", 6, true},
		{"FindValue", 7, true},
		{"FindValue", 8, false},
		{"Remove", 5, false},

		{"Remove", 6, true},
		{"FindValue", 1, false},
		{"FindValue", 2, true},
		{"FindValue", 3, true},
		{"FindValue", 4, true},
		{"FindValue", 5, false},
		{"FindValue", 6, false},
		{"FindValue", 7, true},
		{"FindValue", 8, false},
		{"Remove", 6, false},

		{"Remove", 2, true},
		{"FindValue", 1, false},
		{"FindValue", 2, false},
		{"FindValue", 3, true},
		{"FindValue", 4, true},
		{"FindValue", 5, false},
		{"FindValue", 6, false},
		{"FindValue", 7, true},
		{"FindValue", 8, false},
		{"Remove", 2, false},

		{"Remove", 4, true},
		{"FindValue", 1, false},
		{"FindValue", 2, false},
		{"FindValue", 3, true},
		{"FindValue", 4, false},
		{"FindValue", 5, false},
		{"FindValue", 6, false},
		{"FindValue", 7, true},
		{"FindValue", 8, false},
		{"Remove", 4, false},

		{"Remove", 7, true},
		{"FindValue", 1, false},
		{"FindValue", 2, false},
		{"FindValue", 3, true},
		{"FindValue", 4, false},
		{"FindValue", 5, false},
		{"FindValue", 6, false},
		{"FindValue", 7, false},
		{"FindValue", 8, false},
		{"Remove", 7, false},

		{"Remove", 3, true},
		{"FindValue", 1, false},
		{"FindValue", 2, false},
		{"FindValue", 3, false},
		{"FindValue", 4, false},
		{"FindValue", 5, false},
		{"FindValue", 6, false},
		{"FindValue", 7, false},
		{"FindValue", 8, false},
		{"Remove", 3, false},
	}

	m := NewSet[int](Less[int])

	for _, ts := range tests {
		ts := ts
		title := fmt.Sprintf("%s %d", ts.op, ts.input)
		t.Run("Set:"+title, func(t *testing.T) {
			prev := m.String()
			var got bool
			switch ts.op {
			case "Insert":
				got = m.Insert(ts.input)
			case "FindValue":
				got, _ = m.FindValue(ts.input)
			case "Remove":
				got = m.Remove(ts.input)
			}
			if ts.want != got {
				t.Errorf("Set.%s(%d)=%s failed: want=%v got=%v\n", ts.op, ts.input, prev, ts.want, got)
			}
		})
	}
}

func TestSetInsertRandom(t *testing.T) {
	N := 20

	om := NewSet[int](Less[int])
	m := make(map[int]bool)

	doit := func() {
		prev := om.String()
		key := rand.Int() % N
		var opstr string
		c := rand.Int()
		if c%2 == 0 {
			opstr = fmt.Sprintf("Insert(%d)", key)
			om.Insert(key)
			m[key] = true
		} else {
			opstr = fmt.Sprintf("Remove(%d)", key)
			om.Remove(key)
			delete(m, key)
		}
		post := om.String()
		for k := -1; k <= 20; k++ {
			want := m[k]
			got := om.FindValue(k)
			if want != got {
				t.Fatalf("Set.Find(%d) returns wrong value after %s {%s -> %s}: want=%v got=%v\n", k, opstr, prev, post, want, got)
			}
		}
	} // doit()

	for i := 0; i < 1000; i++ {
		doit()
	}
}

func TestSetLowerBound(t *testing.T) {
	tests := []struct {
		insert int
		want   []int
	}{
		//        1, 2, 3, 4, 5, 6, 7, 8, 9
		{4, []int{4, 4, 4, 4, 0, 0, 0, 0, 0}},
		{2, []int{2, 2, 4, 4, 0, 0, 0, 0, 0}},
		{8, []int{2, 2, 4, 4, 8, 8, 8, 8, 0}},
		{6, []int{2, 2, 4, 4, 6, 6, 8, 8, 0}},
	}

	m := NewSet[int, bool](Less[int])

	for _, ts := range tests {
		ts := ts
		m.Insert(ts.insert)
		title := fmt.Sprintf("after insert %d", ts.insert)
		t.Run("Set.LowerBound:"+title, func(t *testing.T) {
			vs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			for i := range vs {
				var got int
				it, ok := m.LowerBound(vs[i])
				if ok {
					got = it.Key()
				}
				if ts.want[i] != got {
					t.Errorf("Set.LowerBound(%d) on %s failed: want=%v got=%v\n", vs[i], m.String(), ts.want, got)
				}
			}
		})
	}
}

func TestSetUpperBound(t *testing.T) {
	tests := []struct {
		insert int
		want   []int
	}{
		//        1, 2, 3, 4, 5, 6, 7, 8, 9
		{4, []int{4, 4, 4, 0, 0, 0, 0, 0, 0}},
		{2, []int{2, 4, 4, 0, 0, 0, 0, 0, 0}},
		{8, []int{2, 4, 4, 8, 8, 8, 8, 0, 0}},
		{6, []int{2, 4, 4, 6, 6, 8, 8, 0, 0}},
	}

	m := NewSet[int, bool](Less[int])

	for _, ts := range tests {
		ts := ts
		m.Insert(ts.insert)
		title := fmt.Sprintf("after insert %d", ts.insert)
		t.Run("Set.UpperBound:"+title, func(t *testing.T) {
			vs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			for i := range vs {
				var got int
				it, ok := m.UpperBound(vs[i])
				if ok {
					got = it.Key()
				}
				if ts.want[i] != got {
					t.Errorf("Set.UpperBound(%d) on %s failed: want=%v got=%v\n", vs[i], m.String(), ts.want, got)
				}
			}
		})
	}
}
