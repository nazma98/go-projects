package main

import "fmt"

func main() {
	slice := make([]string, 3, 6)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	slice2 := append(slice, "hello", "world")
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(slice[:])
	fmt.Println(len(slice), cap(slice))

}
