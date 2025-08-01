package main

import "fmt"

func main() {
	var p *int
	i := 10
	p = &i

	defer func(x int) {
		fmt.Println(x)
	}(*p)

	i = 20
}
