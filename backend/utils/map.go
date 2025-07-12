package utils

func FromEntries[T any, K comparable](entries []T, keyFn func(T) K) map[K]T {
	m := make(map[K]T, len(entries))
	for _, e := range entries {
		m[keyFn(e)] = e
	}

	return m
}
