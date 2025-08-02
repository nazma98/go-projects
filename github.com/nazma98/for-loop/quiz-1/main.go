package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	for i := 0; i < len(nums); i++ {
		i := nums[i]
		fmt.Print(i, " ")
	}
}
