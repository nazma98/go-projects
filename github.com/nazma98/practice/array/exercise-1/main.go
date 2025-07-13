package main

import "fmt"

func changeArray(arr [3]int) {
	arr[0] = 99
}

func main() {
	a := [3]int{1, 2, 3}
	changeArray(a)
	fmt.Println(a)
}
