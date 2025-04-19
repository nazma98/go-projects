package main

import "fmt"

func main() {
	var task []string
	task = append(task, "Buy Milk")
	task = append(task, "Finish project 1", "Complete videos", "Practice codewars")

	for index, value := range task {
		fmt.Printf("%d. %s\n",index + 1, value)
	}
}