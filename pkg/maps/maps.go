package maps

import "maps"

type Map[K comparable, V any] map[K]V
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func New[K comparable, V any](entries ...Entry[K, V]) Map[K, V] {
	m := make(Map[K, V])
	m.Add(entries...)
	return m
}

func NewEntry[K comparable, V any](key K, value V) Entry[K, V] {
	return Entry[K, V]{key, value}
}

func (m Map[K, V]) Get(key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}

func (m *Map[K, V]) GetOrPanic(key K) *V {
	v, ok := (*m)[key]
	if !ok {
		panic("key not found")
	}
	return &v
}

func (m Map[K, V]) GetOr(key K, defaultValue V) V {
	v, ok := m[key]
	if !ok {
		return defaultValue
	}
	return v
}

func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}

func (m Map[K, V]) Has(key K) bool {
	_, ok := m[key]
	return ok
}

func (m Map[K, V]) Delete(key K) {
	delete(m, key)
}

func (m Map[K, V]) Clear() {
	for k := range m {
		delete(m, k)
	}
}

func (m Map[K, V]) Add(entries ...Entry[K, V]) {
	for _, e := range entries {
		m[e.Key] = e.Value
	}
}

func (m Map[K, V]) AddI(entries ...Entry[K, V]) Map[K, V] {
	n := maps.Clone(m)
	for _, e := range entries {
		n.Add(e)
	}
	return n
}

func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[K, V]) Values() []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func (m Map[K, V]) Entries() []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, Entry[K, V]{k, v})
	}
	return entries
}

func (m Map[K, V]) Map(fn func(k K, v V) (K, V)) Map[K, V] {
	r := make(Map[K, V])
	for k, v := range m {
		k, v = fn(k, v)
		r[k] = v
	}
	return r
}
