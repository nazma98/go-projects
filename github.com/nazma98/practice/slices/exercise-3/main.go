package main

import "fmt"

func main() {
	slice := make([]int, 3)
	fmt.Println(slice)
	slice = make([]int, 4, 7)
	fmt.Println(slice)
}
