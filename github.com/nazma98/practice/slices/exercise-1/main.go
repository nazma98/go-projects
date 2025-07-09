package main

import "fmt"

func main() {
	slice := []int{3, 5, 6, 7}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
