package main

import (
	"fmt"
)

func countSortString(s string) string {
	strSlice := make([]int, 256)

	for i := 0; i < len(s); i++ {
		strSlice[s[i]]++
	}

	sortedStr := make([]byte, 0, len(s))
	for i := 0; i < 256; i++ {
		for strSlice[i] > 0 {
			sortedStr = append(sortedStr, byte(i))
			strSlice[i]--
		}
	}
	return string(sortedStr)
}

func isAnagram(s string, t string) bool {
	if len(s) == 1 && len(t) == 1 {
		if s == t {
			return true
		}
		return false
	}
	str1 := countSortString(s)
	str2 := countSortString(t)

	if str1 == str2 {
		return true
	}
	return false
}

func main() {
	s := "ab"
	t := "ba"
	fmt.Println(isAnagram(s, t))
}
