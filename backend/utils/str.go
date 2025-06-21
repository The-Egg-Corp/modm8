package utils

import "strings"

func ContainsEqualFold(arr []string, item string) bool {
	lowerItem := strings.ToLower(item)

	for _, str := range arr {
		if strings.ToLower(str) == lowerItem {
			return true
		}
	}

	return false
}
