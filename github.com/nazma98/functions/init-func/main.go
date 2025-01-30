package main

import "fmt"

func main() {
	fmt.Println("Hello this is main function")
}

func init() {
	fmt.Println("Hello this is init function which is executed in the beginning")
}