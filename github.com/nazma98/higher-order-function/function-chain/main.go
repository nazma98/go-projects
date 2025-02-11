package main

import "fmt"

func chainFunction(x int, op1 func(a int) int, op2 func(b int) int) int {
	return op1(op2(x))
}

func addOne(x int) int {
	return x + 1
}

func triple(x int) int {
	return x * 3
}

func main() {
	fmt.Println("The final result is after adding one then multiplying with 3 ", chainFunction(5, triple, addOne))
}