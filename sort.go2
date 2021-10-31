package main

import "math"

// vim:set ft=go:

// snip ------------------------------------------------------------------------

const sort_insertion_threashold = 12

// Sorts the given slice according to the given less function using insertion-sort algorithm
func InsertionSort[T any](a []T, less func(T, T) bool) {
	for i := 1; i < len(a); i++ {
		for j := i - 1; 0 <= j && less(a[j+1], a[j]); j-- {
			a[j], a[j+1] = a[j+1], a[j]
		}
	}
}

// Sorts the given slice according to the given less function using heap-sort algorithm
func HeapSort[T any](a []T, less func(T, T) bool) {
	for i := range a {
		for j, p := i, (i-1)/2; j != 0 && less(a[p], a[j]); j, p = p, (p-1)/2 {
			a[j], a[p] = a[p], a[j]
		}
	}
	for i := len(a) - 1; 0 < i; i-- {
		a[0], a[i] = a[i], a[0]
		p := 0
		for {
			l := p*2 + 1 // left child
			r := p*2 + 2 // right child
			if i <= l {
				break
			}
			if i <= r {
				if less(a[p], a[l]) {
					a[p], a[l] = a[l], a[p]
				}
				p = l
				continue
			}

			if less(a[l], a[p]) && less(a[r], a[p]) {
				break
			} else if less(a[l], a[p]) {
				a[p], a[r] = a[r], a[p]
				p = r
			} else if less(a[r], a[p]) {
				a[p], a[l] = a[l], a[p]
				p = l
			} else if less(a[l], a[r]) {
				a[p], a[r] = a[r], a[p]
				p = r
			} else {
				a[p], a[l] = a[l], a[p]
				p = l
			}
		}
	}
}

// Sorts the given slice according to the given less function using quicksort based
// algorithm. Use SortStable for stable sort.
func Sort[T any](a []T, less func(T, T) bool) {
	sort_quick(a, less, intl_CeilPow2(len(a))*2)
}

func sort_quick[T any](a []T, less func(T, T) bool, heap_threshold int) {
	if heap_threshold == 0 {
		HeapSort(a, less)
		return
	}
	if len(a) <= sort_insertion_threashold {
		InsertionSort(a, less)
		return
	}
	sort_quick_median_of_the_three(a, less)
	i := 1
	for i < len(a) && less(a[i], a[0]) {
		i++
	}
	for j := i + 1; j < len(a); j++ {
		if less(a[j], a[0]) {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	i--
	a[0], a[i] = a[i], a[0]
	sort_quick(a[:i], less, heap_threshold-1)
	sort_quick(a[i+1:], less, heap_threshold-1)
}

func sort_quick_median_of_the_three[T any](a []T, less func(T, T) bool) {
	//  Sedgewick'a median of the three
	m := len(a) / 2
	r := len(a) - 1
	a[1], a[m] = a[m], a[1]
	if less(a[r], a[1]) {
		a[1], a[r] = a[r], a[1]
	}
	if less(a[r], a[0]) {
		a[0], a[r] = a[r], a[0]
	}
	if less(a[0], a[1]) {
		a[0], a[1] = a[1], a[0]
	}
}

// Sorts the given slice according to the given less function using mergesort based
// algorithm.
func SortStable[T any](a []T, less func(T, T) bool) {
	buf := make([]T, len(a))
	sort_merge(a, buf, less)
}

func sort_merge[T any](a, buf []T, less func(T, T) bool) {
	if len(a) <= sort_insertion_threashold {
		InsertionSort(a, less)
		return
	}

	m := len(a) / 2
	sort_merge(a[:m], buf, less)
	sort_merge(a[m:], buf, less)

	var i, j, k int
	for i < m && m+j < len(a) {
		if less(a[m+j], a[i]) {
			buf[k] = a[m+j]
			k++
			j++
		} else {
			buf[k] = a[i]
			k++
			i++
		}
	}
	for ; i < m; i++ {
		t := len(a) - (m - i)
		a[i], a[t] = a[t], a[i]
	}
	for i := 0; i < k; i++ {
		a[i] = buf[i]
	}
}

// Sorts the given slice according to the given less function using bubble-sort algorithm
func BubbleSort[T any](a []T, less func(T, T) bool) {
	need := true
	for need {
		need = false
		for j := len(a) - 1; 0 < j; j-- {
			if less(a[j], a[j-1]) {
				a[j], a[j-1] = a[j-1], a[j]
				need = true
			}
		}
	}
}

// Sorts the given slice according to the given less function using selection-sort algorithm
func SelectionSort[T any](a []T, less func(T, T) bool) {
	n := len(a)
	for i := 0; i < n; i++ {
		k := i
		for j := i; j < n; j++ {
			if less(a[j], a[k]) {
				k = j
			}
		}
		a[i], a[k] = a[k], a[i]
	}
}

// Sorts the given slice according to the given less function using shell-sort algorithm
// It uses Tokuda'a sequence for the gaps.
func ShellSort[T any](a []T, less func(T, T) bool) {
	n := len(a)
	G := []int{1}
	for k := 2; ; k++ {
		// Tokuda, Naoyuki (1992). "An Improved Shellsort". In van Leeuven, Jan (ed.).
		// Proceedings of the IFIP 12th World Computer Congress on Algorithms, Software,
		// Architecture. Amsterdam: North-Holland Publishing Co. pp. 449-457.
		// ISBN 978-0-444-89747-3.
		v := int(math.Ceil((math.Pow(float64(9)/4.0, float64(k-1))*9 - 4.0) / 5.0))
		if n <= v {
			break
		}
		G = append(G, v)
	}
	Reverse(G)
	for _, g := range G {
		sort_shell_insertion(a, less, g)
	}
}

func sort_shell_insertion[T any](a []T, less func(T, T) bool, g int) {
	for i := g; i < len(a); i++ {
		tmp := a[i]
		j := i
		for ; g <= j && less(tmp, a[j-g]); j -= g {
			a[j] = a[j-g]
		}
		a[j] = tmp
	}
}

// Sorts the given slice according to the given less function using merge-sort algorithm.
func MergeSort[T any](a []T, less func(T, T) bool) {
	buf := make([]T, len(a))
	sort_merge_naive(a, buf, less)
}

func sort_merge_naive[T any](a, buf []T, less func(T, T) bool) {
	if len(a) == 1 {
		return
	}

	m := len(a) / 2
	sort_merge_naive(a[:m], buf, less)
	sort_merge_naive(a[m:], buf, less)

	var i, j, k int
	for i < m && m+j < len(a) {
		if less(a[m+j], a[i]) {
			buf[k] = a[m+j]
			k++
			j++
		} else {
			buf[k] = a[i]
			k++
			i++
		}
	}
	for ; i < m; i++ {
		t := len(a) - (m - i)
		a[i], a[t] = a[t], a[i]
	}
	for i := 0; i < k; i++ {
		a[i] = buf[i]
	}
}

func CountingSort[T Integer](a []T, less func(T, T) bool) {
	mx := a[0]
	mn := a[0]
	for _, v := range a {
		mn = min(mn, v)
		mx = max(mx, v)
	}
	k := mx + T(1)
	assert(0 <= int(mn) && int(mx) <= 10000000)
	count := make([]int, k)
	for _, v := range a {
		count[v]++
	}
	for i := 1; i < int(k); i++ {
		count[i] += count[i-1]
	}
	buf := make([]T, len(a))
	for _, v := range a {
		buf[count[v]-1] = v
		count[v]--
	}
	copy(a, buf)
}

// Returns true if the given slice is sorted according to the given less function,
// otherwise false,
func IsSorted[T any](a []T, less func(T, T) bool) bool {
	for i := 1; i < len(a); i++ {
		if less(a[i], a[i-1]) {
			return false
		}
	}
	return true
}
