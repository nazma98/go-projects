package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x
	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println(x)
	fmt.Println(y)
}