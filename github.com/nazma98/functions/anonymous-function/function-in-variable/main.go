package main

import "fmt"

func sum() {
	add(2, 4)
}

func add(a, b int) {
	fmt.Println(a + b)
}

func main() {

	add := func (a int, b int) {
		c := a + b
		fmt.Println(c)
	} 

	add(4, 6)
}

func init() {
	fmt.Println("This is init function which excutes in the beginning")
}