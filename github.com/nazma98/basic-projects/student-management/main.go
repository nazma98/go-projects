package main

import (
	"fmt"
	"student-management/student"
)

func main() {
	fmt.Println("📚 Student Management System Started...")
	s1 := student.Student{
		ID:     1,
		Name:   "Nazma",
		Age:    21,
		Email:  "nazma@example.com",
		Grades: []int{84, 64, 79},
	}
	s2 := student.Student{
		ID:     2,
		Name:   "Asha",
		Age:    21,
		Email:  "asha@mail.com",
		Grades: []int{85, 92},
	}
	student.AddStudent(s1)
	student.AddStudent(s2)

	student.PrintAllStudents()

}
