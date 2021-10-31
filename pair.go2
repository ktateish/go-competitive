package main

// vim:set ft=go:

// snip ------------------------------------------------------------------------

type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

func MakePair[T1, T2 any](a T1, b T2) Pair[T1, T2] {
	return Pair[T1, T2]{a, b}
}

type Trio[T1, T2, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

func MakeTrio[T1, T2, T3 any](a T1, b T2, c T3) Trio[T1, T2, T3] {
	return Trio[T1, T2, T3]{a, b, c}
}
