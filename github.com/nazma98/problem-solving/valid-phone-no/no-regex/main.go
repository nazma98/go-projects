package main

import (
	"fmt"
	"strings"
)

func isValidBDPhone(number string) bool {
	if strings.HasPrefix(number, "+8801") {
		if len(number) != 14 {
			return false
		}
		operator := number[5]
		return operator >= '3' && operator <= '9'
	} else if strings.HasPrefix(number, "01") {
		if len(number) != 11 {
			return false
		}
		operator := number[2]
		return operator >= '3' && operator <= '9'
	}
	return false
}

func main() {
	testNumbers := []string{
		"01712345678",    // valid local
		"+8801812345678", // valid intl
		"0112345678",     // invalid (too short)
		"+8801212345678", // invalid operator
		"01987654321",    // valid
		"01312345",       // invalid length
	}

	for _, num := range testNumbers {
		if isValidBDPhone(num) {
			fmt.Println(num, "✅ Valid")
		} else {
			fmt.Println(num, "❌ Invalid")
		}
	}
}
