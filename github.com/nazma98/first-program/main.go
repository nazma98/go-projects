package main

import "fmt"

func main() {
	age := 18

	if age > 18 {
		fmt.Println("You are eligible to marry")
	} else if age < 18 {
		fmt.Println("You are not eligible to marry")
	} else {
		fmt.Println("Oh! Just a teenager")
	}

	b := 30

	switch b {
	case 1:
		fmt.Println("b is 1")
	case 2, 3:
		fmt.Println("b is 2 or 3")
	case 4, 5:
		fmt.Println("b is 4 or 5")
	default:
		fmt.Println("b is neither 1 nor 2 nor 3 nor 4 nor 5")
	}
}
