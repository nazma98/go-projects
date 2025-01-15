package main

import "fmt"

func getNumber() int {
	var num int
	fmt.Println("Enter the number - ")
	fmt.Scanln(&num)
	return num
}

func print(result string) {
	fmt.Println(result)
}

func findOddEven(num int) string {
	if(num % 2 == 1) {
		return "odd"
	}
	return "even"
}

func main() {
	num := getNumber()
	result := findOddEven(num)
	print(result)
}