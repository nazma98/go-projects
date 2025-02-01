package main

import "fmt"

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