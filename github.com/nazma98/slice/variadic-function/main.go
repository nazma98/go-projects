package main

import "fmt"

func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {
	print(4, 5, 7, 8, 9)
}