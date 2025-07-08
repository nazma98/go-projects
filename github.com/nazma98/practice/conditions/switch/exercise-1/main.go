package main

import "fmt"

func main() {
	day := "Friday"

	switch day {
	case "Monday":
		fmt.Println("Start of the Week")
	case "Wednesday":
		fmt.Println("Another day")
	case "Friday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Normal Day")
	}
}
