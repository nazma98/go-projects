package main

import "fmt"

func compareFunctions(x int, op func(a int) int, op2 func(b int) int) int {
	if op(x) > op2(x) {
		return op(x)
	}
	return(op2(x))
}

func double(x int) int {
	return x * 2
}

func square(x int) int {
	return x * x
}

func main(){
	fmt.Println("The larger number is :", compareFunctions(5, double, square))
}