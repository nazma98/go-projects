package main

import "fmt"

func main() {
	slice := []int{1, 3, 5, 7}
	for i := range slice {
		slice[i] *= 2
	}

	fmt.Println(slice)
}
