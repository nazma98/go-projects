package main

import "fmt"

func triangleType(nums []int) string {
	if nums[0] == nums[1] && nums[1] == nums[2] && nums[2] == nums[0] {
		return "equilateral"
	} else if nums[0] == nums[1] && nums[0]+nums[1] > nums[2] ||
		nums[1] == nums[2] && nums[1]+nums[2] > nums[0] ||
		nums[2] == nums[0] && nums[2]+nums[0] > nums[1] {
		return "isosceles"
	} else if nums[0]+nums[1] > nums[2] && nums[2]+nums[1] > nums[0] && nums[0]+nums[2] > nums[1] {
		return "scalene"
	}
	return "none"
}

func main() {
	nums := []int{3, 7, 4}
	fmt.Println(triangleType(nums))
}
