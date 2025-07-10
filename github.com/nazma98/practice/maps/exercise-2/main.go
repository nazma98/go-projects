package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 23,
		"Bob":   25,
	}

	ages["Charlie"] = 34
	fmt.Println(ages)

	delete(ages, "Bob")
	fmt.Println(ages)
}
