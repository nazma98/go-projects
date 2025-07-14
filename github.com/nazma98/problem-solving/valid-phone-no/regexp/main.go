package main

import (
	"fmt"
	"regexp"
)

func isValidBDPhone(number string) bool {
	localPattern := `^01[3-9][0-9]{8}$`
	intPattern := `\+8801[3-9][0-9]{8}$`

	localMatch, _ := regexp.MatchString(localPattern, number)
	intlMatch, _ := regexp.MatchString(intPattern, number)

	return localMatch || intlMatch
}

func main() {
	testNumbers := []string{
		"01712345678",    // valid local
		"+8801712345678", // valid intl
		"0112345678",     // invalid
		"+8801012345678", // invalid operator prefix
		"01712345",       // invalid length
	}
	for _, num := range testNumbers {
		if isValidBDPhone(num) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
