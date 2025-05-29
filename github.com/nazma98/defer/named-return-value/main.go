package main

import "fmt"

func add(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	res := add(4, 5)
	fmt.Println(res)
}
