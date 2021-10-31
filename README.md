# Go Library for Competitive Programming with Generics

Go used to be a difficult language to use for competitive programming. However,
with the introduction of generics, Go has become a language with great
potential for competitive programming.

This repository is a library of the Go language that is useful for competitive
programming.  The code in this library does not compile with the current
official release of Go 1.17.x, but it will compile with Go 1.17.x if you
convert it with [go2go](https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md).
Furthermore, by using [Gottani](https://github.com/ktateish/gottani) to compile
it into a single go file, it is possible to generate Go code that can be
submitted to online judges for competitive programming. The author of this
repository has actually competed in many competitive programming rounds with
Go code generated using this library.


## Supported algorithms

### [AC Library](https://github.com/atcoder/ac-library) Porting

* Data structure
    * [x] FenwickTree
    * [x] Segtree
    * [x] LazySegtree
    * [x] SuffixArray
    * [x] LCPArray
    * [x] ZAlgorithm
* Mathematic
    * [x] PowMod
    * [x] InvMod
    * [x] CRT
    * [x] FloorSum
    * [ ] Convolution, not ported yet.
    * [x] ModInt
* Graph
    * [x] DSU
    * [x] MaxFlow
    * [x] MinCostFlow
    * [x] SCC
    * [x] TwoSat

### Other Data Structure and Algorithms

* Data Structure
    * Deque
    * Heap
        *HeapPush
        *HeapPop
    * LCS: Longest Common Subsequence
    * LCSR: Longest Common Subsequence with Reconsturuction
    * LIS: Longest Incresing Subsequence
    * List: Doubly Linked List
    * LLRB: Left-Leaning Red Brack Tree
    * Map: Ordered Map using LLRB
    * Set: Ordered Set using LLRB
    * Pair
    * Trio
    * Permutation
        * NextPermutation
        * PrevPermutation
    * PriorityQueue
    * Queue
    * Rolling Hash
        * StringContains
        * StringFindAll
        * SequenceContains
        * SequenceFindAll
        * SequenceFindAll2D
    * SkewHeap
    * Stack
    * Treap
    * Trie
* Algorithms for slice []T
    * CountInversion: Find the numbers of pairs (i, j) where a[i] > a[j]
    * Binary Search
        * BinSearch
        * LowerBound
        * UpperBound
        * EqualRange
        * EqualRange
    * Reverse
    * Rotate
    * Shuffle
    * SliceMaxElement
    * SliceFilter
    * SliceMap
    * SliceFoldl
    * SliceFoldr
    * SliceLess
    * SliceUniq
    * Iota
    * SliceBinSearch
    * SliceLowerBound
    * SliceUpperBound
    * SliceEqualRange
    * SliceToString
    * SliceFill
    * SliceEqual
* Sort
    * Sort: Introsort
    * SortStable: Stable sort by Merge Sort with Insertion Sort
    * InsertionSort
    * HeapSort
    * BubbleSort
    * SelectionSort
    * ShellSort
    * MergeSort
    * CountingSort
    * IsSorted
* Algorithms for grid [][]T
    * CumulativeSum2D
* Mathematic
    * IsPrime
    * Primes: Compute prime / composite for range 1 .. n
    * PrimesX: Same as Primes but it can find prime factors.
    * Factorize
    * NewFactorial: Compute combination, permutation, factorial with p as the law
* Graph
    * Dijkstra
    * GetParent: Reconstruct parents in the minimum cost path with the out put of the Dijkstra
    * BellmanFord
    * WarshallFloyd
    * MST: Minimum Spanning Tree
* Utilities
    * GetInf: Get numeric value whitch is useful as "infinit" in competitive programming
    * Less
    * More
    * GetRandomSeed
* IO
    * readb: Reads next string and returns as []byte
    * reads: Reads next string and returns as string
    * readbln: Reads next line and returns as []byte
    * readsln: Reads next line and returns as string
    * readll: Reads next string as int64
    * readi: Reads next string as int
    * readf: Reads next string as float64
    * readInts: Reads next n ints and returns as []int
    * printf: fmt.Printf
    * println: fmt.Println
    * eprintf: fmt.Frintf(os.Stderr, "...")
    * eprintln: fmt.Frintln(os.Stderr, "...")
    * dbgf: fmt.Frintf(os.Stderr, "...")
    * dbg: fmt.Frintln(os.Stderr, "...")
    * sprintf: fmt.Sprintf
    * sprint: fmt.Sprint


## Install Prerequisites

* Install [go2go](https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md).
* Install [Gottani](https://github.com/ktateish/gottani).


## Tutorial

Let's use this lib for solving [ABC 183 C - Travel](https://atcoder.jp/contests/abc183/tasks/abc183_c).

1. Clone [this repository](https://github.com/ktateish/go-competitive).

```
$ git clone https://github.com/ktateish/go-competitive abc183_c
$ cd abc183_c
```

2. Edit main.go2 like:

```go
package main

// vim:set ft=go:

// snip ------------------------------------------------------------------------

const (
	// big prime
	P1b7 = 1000000007
	P1b9 = 1000000009
	P0b3 = 998244353
)

var DEBUG bool
var RANDOM_SEED int64

func Main() {
	N := readi()
	K := readi()
	T := make([][]int, N)
	for i := range T {
		_, T[i] = readInts(N)
	}
	M := N - 1
	a := make([]int, M)
	Iota(a, 1)
	var ans int
	for ok := true; ok; ok = NextPermutation(a) {
		var k int
		k += T[0][a[0]]
		k += T[a[M-1]][0]
		for i := 0; i < M-1; i++ {
			k += T[a[i]][a[i+1]]
		}
		if k == K {
			ans++
		}
	}
	println(ans)
}
```

3. Build using `make`

```
$ make
```

A single Go file `gottani.go` and an executable `a.out` will be generated.

4. Test `a.out` for sample input

```
$ ./a.out < input
```

5. Submit `gottani.go` to the judge and check the result

In this case, the submission is [accepted](https://atcoder.jp/contests/abc183/submissions/26955611)

DISCLAIMER: The author and the contributors of this library are not responsible
for any damages caused by using this library.


## License

This library is distribted under the BSD License (BSD-2-Clause).
See LICENSE file for detail.


## Author

Katsuyuki Tateish (@ktateish)

