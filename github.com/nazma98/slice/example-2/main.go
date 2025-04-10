package main

import "fmt"

func main() {
	x := [] int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]
	
	fmt.Println(x)
	fmt.Println(a)
}