package main

import "fmt"

func main() {
	age := 20

	if age > 18 {
		fmt.Println("You are eligible to marry")
	} else if age < 18 {
		fmt.Println("You are not eligible to marry")
	} else {
		fmt.Println("Oh! Just a teenager")
	}
}
