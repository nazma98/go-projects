package main

import "fmt"
import "example.com/mathlib"

func main() {
	fmt.Println("Showing custom package")

	mathlib.Add(4, 7)
}