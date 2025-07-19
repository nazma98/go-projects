package main

import "fmt"

type Employee struct {
	Name   string
	Id     string
	Salary float64
	Title  string
}

func (e *Employee) AnnualSalary() float64 {
	return 12 * e.Salary
}

func (e *Employee) Promote(title string, increment float64) {
	e.Title = title
	e.Salary += increment
}

func (e *Employee) Info() {
	fmt.Printf("%s with ID %s is a %s with Annual Salary %f", e.Name, e.Id, e.Title, e.AnnualSalary())
}

func main() {
	employee1 := Employee{
		Name:   "Alice Rahman",
		Id:     "E2343",
		Salary: 50000,
		Title:  "Software Engineer",
	}
	employee1.Info()
	employee1.Promote("Senior Software Engineer", 15000)
	employee1.Info()

}
