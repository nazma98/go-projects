package main

import "fmt"

func main() {
	slice := make([]string, 3, 6)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	slice2 := append(slice, "hello", "world")
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(slice[:])
	fmt.Println(len(slice), cap(slice))
	// Print addresses of slice headers (slice variables themselves)
	fmt.Printf("Address of slice header : %p\n", &slice)
	fmt.Printf("Address of slice2 header: %p\n", &slice2)

	// Print addresses of underlying data
	fmt.Printf("Address of slice[0]     : %p\n", &slice[0])
	fmt.Printf("Address of slice2[0]    : %p\n", &slice2[0])
}
