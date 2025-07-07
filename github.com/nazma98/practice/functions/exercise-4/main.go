package main

import "fmt"

func getCoords() (x, y int) {
	x = 4
	y = 8
	return
}

func main() {
	x, y := getCoords()
	fmt.Println(x, ",", y)
}
