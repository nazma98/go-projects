package main

import "fmt"

func calculate(a, b int, op func(x, y int) int) int {
	return op(a, b)
}

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func multiply(a, b int) int {
    return a * b
}

func main(){
		fmt.Println("The sum is ", calculate(6, 4, add))
		fmt.Println("The subtraction is ",calculate(12, 7, subtract))
		fmt.Println("The product is ",calculate(6, 4, multiply))
}