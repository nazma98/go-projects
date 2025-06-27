package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		total := nums[left] + nums[right]
		if total == target {
			break
		} else if total > target {
			right -= 1
		} else {
			left += 1
		}
	}
	return []int{left + 1, right + 1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	fmt.Println(ans)
}
