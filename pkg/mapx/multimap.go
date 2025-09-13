package mapx

type MultiMap[K comparable, V any] map[K][]V

func NewMultiMap[K comparable, V any]() MultiMap[K, V] {
	return make(MultiMap[K, V])
}

func (m MultiMap[K, V]) Add(key K, value V) {
	m[key] = append(m[key], value)
}

func (m MultiMap[K, V]) Get(key K) []V {
	return m[key]
}
