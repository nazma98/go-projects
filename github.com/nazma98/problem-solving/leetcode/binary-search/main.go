package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int

	for left <= right {
		mid = (right-left)/2 + left

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	nums := []int{2, 4, 6, 8, 23, 45, 67}
	fmt.Println(search(nums, 22))
}
