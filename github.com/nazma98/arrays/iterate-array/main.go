package main

import "fmt"

func main() {
	arr := [...] int {2, 4, 7, 9, 12, 34}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for index, value := range arr {
		fmt.Printf("Index %d value %d\n", index, value)
	}
}