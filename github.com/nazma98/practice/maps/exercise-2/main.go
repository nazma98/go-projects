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

	val, ok := ages["Bob"]
	if ok {
		fmt.Println("Bob Exists", val)
	} else {
		fmt.Println("Bob doesn't exist")
	}
}
