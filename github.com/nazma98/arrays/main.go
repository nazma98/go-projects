package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 4}
	arr3 := [...]int{2, 5, 7, 90}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}