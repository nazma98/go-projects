package main

import "fmt"

func factorial(num int) int {
	result := 1
	if num == 0 || num == 1 {
		return result
	}
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}

func main() {
	ans := factorial(5)
	fmt.Println(ans)
}
