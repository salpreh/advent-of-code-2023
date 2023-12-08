package utils

func AppendAndShift[V any](arr []V, pos int, item V) {
	if len(arr) < 2 {
		(arr)[pos] = item
		return
	}

	for i := len(arr) - 1; i > pos; i-- {
		(arr)[i] = (arr)[i-1]
	}
	(arr)[pos] = item
}

func GetOrDefault[M ~map[K]V, K comparable, V any](m M, key K, defaultValue V) V {
	if v, exists := m[key]; exists {
		return v
	}

	return defaultValue
}

func CountByItem[V comparable](items []V) map[V]int {
	countByItem := make(map[V]int)
	for _, card := range items {
		count := GetOrDefault(countByItem, card, 0)
		count += 1
		countByItem[card] = count
	}

	return countByItem
}

func ToMap[K comparable, V any](items []V, keyExtractor func(V) K) map[K]V {
	result := make(map[K]V)
	for _, it := range items {
		result[keyExtractor(it)] = it
	}

	return result
}

func ToMapP[K comparable, V any](items []V, keyExtractor func(V) K) map[K]*V {
	result := make(map[K]*V)
	for i, _ := range items {
		it := items[i]
		result[keyExtractor(it)] = &it
	}

	return result
}

func Sum(items []int) int {
	sum := 0
	for _, it := range items {
		sum += it
	}

	return sum
}

type Set[T comparable] struct {
	data map[T]bool
}

func NewEmptySet[T comparable]() *Set[T] {
	return &Set[T]{make(map[T]bool)}
}

func NewSet[T comparable](items []T) *Set[T] {
	data := make(map[T]bool)
	for _, item := range items {
		data[item] = true
	}

	return &Set[T]{data}
}

func (s *Set[T]) Add(item T) {
	s.data[item] = true
}

func (s *Set[T]) Remove(item T) {
	delete(s.data, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, exists := s.data[item]

	return exists
}
