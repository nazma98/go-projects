package main

import "fmt"

var sum = 10

func main() {
	fmt.Println(sum)

	sum := func(a, b int) {
		fmt.Println(a + b)
	}

	sum(3, 9)
}