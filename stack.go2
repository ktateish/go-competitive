package main

// vim:set ft=go:

// snip ------------------------------------------------------------------------

type Stack[T any] struct {
	a []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func (st *Stack[T]) Len() int {
	return len(st.a)
}

func (st *Stack[T]) Empty() bool {
	return st.Len() == 0
}

func (st *Stack[T]) Top() T {
	return st.a[len(st.a)-1]
}

func (st *Stack[T]) Push(x T) {
	st.a = append(st.a, x)
}

func (st *Stack[T]) Pop() T {
	n := len(st.a) - 1
	x := st.a[n]
	st.a = st.a[:n]
	return x
}

func (st *Stack[T]) Clear() {
	st.a = st.a[:0]
}

type RTStackNode[T any] struct {
	val  T
	next *RTStackNode[T]
}

func NewRTStack[T any]() *RTStackNode[T] {
	var res *RTStackNode[T]
	return res
}

func (st *RTStackNode[T]) Empty() bool {
	return st == nil
}

func (st *RTStackNode[T]) Top() T {
	return st.val
}

func (st *RTStackNode[T]) Push(x T) *RTStackNode[T] {
	return &RTStackNode[T]{
		val:  x,
		next: st,
	}
}

func (st *RTStackNode[T]) Pop() *RTStackNode[T] {
	return st.next
}
