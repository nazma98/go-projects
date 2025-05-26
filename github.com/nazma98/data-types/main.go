package main

import "fmt"

func main() {
	var a int8 = -128
	var b int8 = 127

	var x uint8 = 255

	var j float32 = 10.2423
	var k float64 = 10.4213

	// bool is stored as 1 byte / 8 bit
	var flag bool = false

	heart := '♠'

	fmt.Printf("%c\n", heart)
	fmt.Printf("%f\n", j)
	fmt.Printf("%.2f\n", k)

	fmt.Println(a, b, x, j, k, flag)
}
