package main

import "fmt"

func main() {
	a := 1
	defer fmt.Println(a) // evaluated now
	a = 2
}
