package main

import "fmt"

func main() {
	slice := make([]int, 4)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	slice2 := append(slice, 2)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))
}
