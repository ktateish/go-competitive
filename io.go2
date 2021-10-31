package main

// vim:set ft=go:

import (
	"fmt"
	"os"
)

// snip ------------------------------------------------------------------------

var bio *BufferedIO

// IO functions

// Reads next token and return it as []byte
func readb() []byte {
	b, err := bio.ReadBytes()
	if err != nil {
		panic(err)
	}
	return b
}

// Reads next token and return it as string
func reads() string {
	return string(readb())
}

// Read next line as []byte.  Trailing '\n' will not be included.
// See also comments on readb()
func readbln() []byte {
	b, err := bio.ReadLine()
	if err != nil {
		panic(err)
	}
	return b
}

// Read next line as string
func readsln() string {
	return string(readbln())
}

// Reads next token and return it as int64
func readll() int64 {
	i, err := bio.ReadInt64()
	if err != nil {
		panic(err.Error())
	}
	return i
}

// Reads next token and return it as int
func readi() int {
	return int(readll())
}

// Reads next token and return it as float64
func readf() float64 {
	f, err := bio.ReadFloat64()
	if err != nil {
		panic(err.Error())
	}
	return f
}

// Write args to bio  with the format f
func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(bio, f, args...)
}

// Write args to bio  without format
func println(args ...interface{}) (int, error) {
	return fmt.Fprintln(bio, args...)
}

// Write args to stderr with the format f
func eprintf(f string, args ...interface{}) {
	if !DEBUG {
		return
	}
	fmt.Fprintf(os.Stderr, f, args...)
}

// Write args to stderr without format
func eprintln(args ...interface{}) {
	if !DEBUG {
		return
	}
	fmt.Fprintln(os.Stderr, args...)
}

func dbgf(f string, args ...interface{}) {
	if !DEBUG {
		return
	}
	fmt.Fprintf(os.Stderr, f, args...)
}

func dbg(args ...interface{}) {
	if !DEBUG {
		return
	}
	fmt.Fprintln(os.Stderr, args...)
}

func sprintf(f string, args ...interface{}) string {
	return fmt.Sprintf(f, args...)
}

func sprint(args ...interface{}) string {
	return fmt.Sprint(args...)
}

// -----------------------------------------------------------------------------

// Utilities

func pyesno(t bool) {
	if t {
		println("Yes")
	} else {
		println("No")
	}
}

func readInts(n int) (int, []int) {
	if n == 0 {
		n = readi()
	}
	a := make([]int, n)
	for i := range a {
		a[i] = readi()
	}
	return n, a
}
