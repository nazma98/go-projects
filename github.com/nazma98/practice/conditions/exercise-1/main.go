package main

import "fmt"

func checkNumber(num int) string {

	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

func main() {
	num := 0
	result := checkNumber(num)
	fmt.Println(result)
}
