package main

import (
	"fmt"
	"student-management/student"
)

func main() {
	fmt.Println("📚 Student Management System Started...")
	s := student.Student{
		ID:     1,
		Name:   "Nazma",
		Age:    21,
		Email:  "nazma@example.com",
		Grades: []int{84, 64, 79},
	}
	s.DisplayInfo()
}
