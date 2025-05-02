// https://leetcode.com/problems/contains-duplicate/description/
package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{4, 1, 7, 32, 75, 7}
	fmt.Println(containsDuplicate(nums))
}
