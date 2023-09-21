package Lecture1

import (
	"sort"
)

func compareTwoSlicesWithOrder(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func compareTwoSlicesWithoutOrder(s1 []int, s2 []int) bool {
	sort.Ints(s1)
	sort.Ints(s2)
	if len(s1) != len(s2) {
		return false
	}
	return compareTwoSlicesWithOrder(s1, s2)   
}
