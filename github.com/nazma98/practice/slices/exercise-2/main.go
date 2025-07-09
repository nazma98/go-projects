package main

import "fmt"

func main() {
	nums := []int{4, 6, 8, 9}

	nums = append(nums, 3, 7)

	fmt.Println(nums)
}
