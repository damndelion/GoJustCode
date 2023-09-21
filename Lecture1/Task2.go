package Lecture1

import "strings"

func longestCommonPrefix(s []string) string {
	prefix := s[0]
	for i := 1; i < len(s); i++ {
		for !strings.HasPrefix(s[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
