package main

import "fmt"

func a() {
	i := 0

	fmt.Println("first", i)

	defer fmt.Println("second", i)

	i = i + 1

	fmt.Println("third", i)

	return
}

func main() {
	a()
}
