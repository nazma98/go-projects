package main

import "fmt"

func findOddEven(num int) string {
	if(num % 2 == 1) {
		return "odd"
	}
	return "even"
}

func main() {
	var num int
	fmt.Println("Enter the number - ")
	fmt.Scanln(&num)
	fmt.Println(findOddEven(num))
}