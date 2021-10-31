package main

// vim:set ft=go:

// snip ------------------------------------------------------------------------

type Trie[V any] struct {
	child [256]*Trie[V]
	val   V
	term  bool
}

func NewTrie[V any]() *Trie[V] {
	return &Trie[V]{}
}

func (t *Trie[V]) Put(key string, val V) {
	t.BPut([]byte(key), val)
}

func (t *Trie[V]) BPut(key []byte, val V) {
	u := t
	for i := 0; i < len(key); i++ {
		k0 := int(key[i])
		next := u.child[k0]
		if next == nil {
			next = &Trie[V]{}
			u.child[k0] = next
		}
		u = next
	}
	u.term = true
	u.val = val
}

func (t *Trie[V]) Get(key string) (V, bool) {
	return t.BGet([]byte(key))
}

func (t *Trie[V]) BGet(key []byte) (V, bool) {
	u := t
	for i := 0; i < len(key); i++ {
		k0 := int(key[i])
		next := u.child[k0]
		if next == nil {
			var res V
			return res, false
		}
		u = next
	}
	return u.val, u.term
}
