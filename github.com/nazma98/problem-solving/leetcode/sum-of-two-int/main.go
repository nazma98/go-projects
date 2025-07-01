package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		sumAnd := a & b
		a = a ^ b
		b = sumAnd << 1
	}
	return a
}

func main() {
	fmt.Println(getSum(4, 8))
}
