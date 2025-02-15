package main

import "fmt"

func getNumber() int {
	fmt.Println("Enter the starting number : ")
	var startNumber int
	fmt.Scanln(&startNumber)
	return startNumber
}

func getIncrementValue() int {
	fmt.Println("Enter the increment value : ")
	var incrementValue int
	fmt.Scanln(&incrementValue)
	return incrementValue
}

func getNumberOfTimes() int {
	fmt.Println("Enter the number of times to increase : ")
	var noOfTimes int
	fmt.Scanln(&noOfTimes)
	return noOfTimes
}

func increaseNumber(startNumber, increment, noOfTimes int) int {
	for i := 1; i <= noOfTimes; i++ {
		startNumber = startNumber + increment
	}
	return startNumber
}

func main() {
	startNumber := getNumber()
	increment := getIncrementValue()
	noOfTimes := getNumberOfTimes()
	result := increaseNumber(startNumber, increment, noOfTimes)
	fmt.Println(result)
}