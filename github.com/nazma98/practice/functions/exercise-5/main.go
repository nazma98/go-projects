package main

import "fmt"

func addNumbers(nums ...int) (sum int) {
	for _, val := range nums {
		sum += val
	}
	return
}

func main() {
	sum := addNumbers(1, 4, 2, 6, 8)
	fmt.Println(sum)
}
