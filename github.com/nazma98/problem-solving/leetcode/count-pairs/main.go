package main

import (
	"fmt"
)

func countPairs(nums []int, target int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] < target {
			fmt.Println(i, " i ", j, " j", nums[i]+nums[j], " sum")
			count++
		}
	}
	return count
}

func main() {
	nums := []int{-6, 2, 5, -2, -7, -1, 3}
	//    0  1  2   3   4    5  6
	target := -2
	fmt.Println(countPairs(nums, target))
}
