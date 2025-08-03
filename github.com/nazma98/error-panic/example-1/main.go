package main

import (
	"errors"
	"fmt"
)

// 1. Returns error if input is invalid
func calculateAge(days int) (int, error) {
	if days <= 0 {
		return 0, errors.New("days must be positive")
	}
	age := days / 365
	return age, nil
}

// 2. Panics if pointer is nil (programmer bug)
func getDaysFromPointer(ptr *int) int {
	if ptr == nil {
		panic("input pointer is nil") // panic: serious problem
	}
	return *ptr
}

func main() {
	// ✅ Example of handling error
	age, err := calculateAge(0) // passing invalid days
	if err != nil {
		fmt.Println("Error:", err) // still continues
	} else {
		fmt.Println("Age is:", age)
	}

	// ❌ Example of panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	var ptr *int = nil
	days := getDaysFromPointer(ptr) // this will panic
	fmt.Println("Got days:", days)

	fmt.Println("Program finished.")
}
