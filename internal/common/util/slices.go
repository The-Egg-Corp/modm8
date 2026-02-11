package util

import "strings"

func FromEntries[T any, K comparable](entries []T, keyFn func(T) K) map[K]T {
	m := make(map[K]T, len(entries))
	for _, e := range entries {
		m[keyFn(e)] = e
	}

	return m
}

// Reports whether the item is contained within s.
// Ignores case-sensitivity - both item and elements of slice s are lower cased automatically.
// Similar to [strings.EqualFold], but takes an a slice of strings instead of a single string.
func SliceEqualFold(s []string, item string) bool {
	for _, str := range s {
		if strings.EqualFold(str, item) {
			return true
		}
	}

	return false
}
