package main

import (
	"fmt"
)

func main() {
	s := "hello"
	// s[0] = 'H' // not allowed
	b := []byte(s)
	b[0] = 'H'
	fmt.Println(b)
	fmt.Println(string(b))
}
