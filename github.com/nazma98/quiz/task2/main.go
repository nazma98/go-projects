package main

import "fmt"

func rectangleMethods(length, width float64, op func(x, y float64) float64) float64 {
	return op(length, width)
}

func rectangleArea(length, width float64) float64 {
    return length * width
}

func rectanglePerimeter(length, width float64) float64 {
    return 2 * (length + width)
}

func main() {
    length := 10.5
    width := 5.2

    area := rectangleMethods(length, width, rectangleArea)
    perimeter := rectangleMethods(length, width, rectanglePerimeter)

    fmt.Println("Rectangle Area:", area)
    fmt.Println("Rectangle Perimeter:", perimeter)
}