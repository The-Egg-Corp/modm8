package utils

import "strings"

// Reports whether the item is contained with. Ignores case-sensitivity - both item and arr elements are lower cased automatically.
//
// Like [strings.EqualFold], however this function takes an an array of strings instead of a single string.
func ArrEqualFold(arr []string, item string) bool {
	lowerItem := strings.ToLower(item)
	for _, str := range arr {
		if strings.ToLower(str) == lowerItem {
			return true
		}
	}

	return false
}
