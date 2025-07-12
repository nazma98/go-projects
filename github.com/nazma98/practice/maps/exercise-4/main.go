package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 3, 1, 4}
	seen := make(map[int]bool)
	unique := []int{}

	for _, val := range nums {
		if !seen[val] {
			seen[val] = true
			unique = append(unique, val)
		}
	}

	fmt.Println(unique)
}
