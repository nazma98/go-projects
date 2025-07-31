package main

import "fmt"

func main() {
	// arr := [4]int{1, 2, 3, 4}
	arr := []int{1, 2, 3, 4}
	slice := arr[3:3] //2 3
	fmt.Println(slice)
	slice = append(slice, 7)
	slice = append(slice, 8)
	// slice = append(slice, 9, 10)
	slice[0] = 10
	fmt.Println("len:", len(arr), "cap:", cap(arr))
	fmt.Println("len:", len(slice), "cap:", cap(slice))

	fmt.Println(arr)
	fmt.Println(slice)

	fmt.Printf("arr[3] addr:   %p\n", &arr[3])
	fmt.Printf("slice[0] addr: %p\n", &slice[0])
}
