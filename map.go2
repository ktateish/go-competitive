package main

// vim:set ft=go:

// snip ------------------------------------------------------------------------

// Map replesents a ordered map using LLRB
type Map[K, V any] struct {
	t *LLRB[K, V]
}

// NewMap returns the pointer a Map object
func NewMap[K, V any](less func(a, b K) bool) *Map[K, V] {
	return &Map[K, V]{
		t: NewLLRB[K, V](less),
	}
}

// Len returns the size of the Map
func (m *Map[K, V]) Len() int {
	return m.t.Len()
}

// Insert inserts the key
func (m *Map[K, V]) Insert(key K, val V) bool {
	return m.t.Insert(key, val)
}

// Remove removes the key
func (m *Map[K, V]) Remove(key K) bool {
	return m.t.Remove(key)
}

// Find returns the MapIterator that matches the key.  The second value is nil if the key is not found.
func (m *Map[K, V]) Find(key K) (MapIterator[K, V], bool) {
	it, ok := m.t.Find(key)
	return MapIterator[K, V](it), ok
}

// Find returns the value of the key.  The second value is nil if the key is not found.
func (m *Map[K, V]) FindValue(key K) (V, bool) {
	return m.t.FindValue(key)
}

func (m *Map[K, V]) Begin() (MapIterator[K, V], bool) {
	it, ok := m.t.Begin()
	return MapIterator[K, V](it), ok
}

func (m *Map[K, V]) End() MapIterator[K, V] {
	it := m.t.End()
	return MapIterator[K, V](it)
}

func (m *Map[K, V]) RBegin() (MapReverseIterator[K, V], bool) {
	it, ok := m.t.RBegin()
	return MapReverseIterator[K, V](it), ok
}

func (m *Map[K, V]) REnd() MapReverseIterator[K, V] {
	it := m.t.REnd()
	return MapReverseIterator[K, V](it)
}

// LowerBound returns the MapIterator that has the lowest key which is greater than or equal to the specified key.
func (m *Map[K, V]) LowerBound(key K) (MapIterator[K, V], bool) {
	it, ok := m.t.LowerBound(key)
	return MapIterator[K, V](it), ok
}

// UpperBound returns the MapIterator that has a key which is lesser than or equal to the specified key.
func (m *Map[K, V]) UpperBound(key K) (MapIterator[K, V], bool) {
	it, ok := m.t.UpperBound(key)
	return MapIterator[K, V](it), ok
}

func (m *Map[K, V]) EqualRange(key K) (MapIterator[K, V], MapIterator[K, V]) {
	lo, hi := m.t.EqualRange(key)
	return MapIterator[K, V](lo), MapIterator[K, V](hi)
}

// String returns the string representation of the Map
func (m *Map[K, V]) String() string {
	return m.t.String()
}

// Graphviz returns the graph directives of the Map in graphviz format
func (m *Map[K, V]) Graphviz() string {
	return m.t.Graphviz()
}

//
// Map Iterator
//

type MapIterator[K, V any] LLRBIterator[K, V]

func (it MapIterator[K, V]) IsEnd() bool {
	return LLRBIterator[K, V](it).IsEnd()
}

func (it MapIterator[K, V]) Key() K {
	return LLRBIterator[K, V](it).Key()
}

func (it MapIterator[K, V]) Value() V {
	return LLRBIterator[K, V](it).Value()
}

func (it MapIterator[K, V]) String() string {
	return LLRBIterator[K, V](it).String()
}

func (it MapIterator[K, V]) Next() (MapIterator[K, V], bool) {
	next, ok := LLRBIterator[K, V](it).Next()
	return MapIterator[K, V](next), ok
}

func (it MapIterator[K, V]) Prev() (MapIterator[K, V], bool) {
	prev, ok := LLRBIterator[K, V](it).Prev()
	return MapIterator[K, V](prev), ok
}

func (it MapIterator[K, V]) Less(jt MapIterator[K, V]) bool {
	return LLRBIterator[K, V](it).Less(LLRBIterator[K, V](jt))
}

type MapReverseIterator[K, V any] LLRBReverseIterator[K, V]

func (it MapReverseIterator[K, V]) IsEnd() bool {
	return LLRBReverseIterator[K, V](it).IsEnd()
}

func (it MapReverseIterator[K, V]) Key() K {
	return LLRBReverseIterator[K, V](it).Key()
}

func (it MapReverseIterator[K, V]) Value() V {
	return LLRBReverseIterator[K, V](it).Value()
}

func (it MapReverseIterator[K, V]) String() string {
	return LLRBReverseIterator[K, V](it).String()
}

func (it MapReverseIterator[K, V]) Next() (MapReverseIterator[K, V], bool) {
	next, ok := LLRBReverseIterator[K, V](it).Next()
	return MapReverseIterator[K, V](next), ok
}

func (it MapReverseIterator[K, V]) Prev() (MapReverseIterator[K, V], bool) {
	prev, ok := LLRBReverseIterator[K, V](it).Prev()
	return MapReverseIterator[K, V](prev), ok
}

func (it MapReverseIterator[K, V]) Less(jt MapReverseIterator[K, V]) bool {
	return LLRBReverseIterator[K, V](it).Less(LLRBReverseIterator[K, V](jt))
}
