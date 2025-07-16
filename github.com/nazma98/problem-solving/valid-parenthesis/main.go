package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if bracketMap[char] != top {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("({[]})")) // true
	fmt.Println(isValid("((])"))   // false
}
