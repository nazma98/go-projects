package main

import "fmt"

func main() {
	str := "Hello"
	freq := make(map[rune]int)

	for _, ch := range str {
		freq[ch]++
	}

	for k, v := range freq {
		fmt.Printf("%c: %d\n", k, v)
	}
}
