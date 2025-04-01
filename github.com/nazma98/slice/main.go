package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "Go", "Interview", "question"}
	fmt.Println(arr)

	s := arr[1:4]
	fmt.Println(s)
	fmt.Println(len(s), " ", cap(s))
}