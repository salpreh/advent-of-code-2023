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
