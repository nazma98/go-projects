package main

import "fmt"

func getDuplicateArray(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(getDuplicateArray(nums))
}
