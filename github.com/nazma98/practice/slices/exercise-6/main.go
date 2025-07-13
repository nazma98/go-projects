package main

import "fmt"

func changeSlice(sl []int) {
	sl[0] = 99
}

func main() {
	sl := []int{2, 5, 7}
	changeSlice(sl)
	fmt.Println(sl)
}
