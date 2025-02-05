package main

import "fmt"

func main() {
	arr := [...]int {2, 5, 9, 23, 54}
	fmt.Println(arr[4])

	arr[4] = 900
	fmt.Println(arr[4])
	fmt.Println(arr)
}