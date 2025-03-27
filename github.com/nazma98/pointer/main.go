package main

import "fmt"

type User struct{
	Name string
	Age int
	Salary float64
}

func print(numbers *[3]int) {
	fmt.Println(numbers)
}

func main() {
	a := 20

	p := &a
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	arr := [3]int{1, 3, 5}
	print(&arr)

	obj := User{
		Name: "Shimul",
		Age: 13,
		Salary: 0,
	}

	fmt.Println(obj)
}