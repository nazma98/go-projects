package main

import "fmt"

func changeSlice(a []int) []int {
	a[0] = 10
	a = append(a, 8)
	return a
}

func main() {
	x := [] int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)
	
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(y)

	fmt.Println(x[0:8])
}