package main

import "fmt"

const a = 10

var (
	p = 100
)

func outer() func() {
	money := 100
	age := 30
	fmt.Println("Age ", age)
	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func main() {

}