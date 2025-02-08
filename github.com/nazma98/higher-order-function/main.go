package main

import "fmt"

func processOperation(a int, b int, op func (x int, y int)) {
	op(a, b)
}

func add(p int, q int) {
	r := p + q
	fmt.Println(r)
}