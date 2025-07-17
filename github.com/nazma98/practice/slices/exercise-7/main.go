package main

import "fmt"

// func main() {
// 	x := make([]int, 0, 2)
// 	x = append(x, 1)
// 	y := x
// 	x = append(x, 2)
// 	y = append(y, 3)

// 	fmt.Println(x)
// 	fmt.Println(y)
// }

// func main() {
// 	a := []int{1, 2, 3}
// 	b := a[:2]
// 	c := append(b, 10)
// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// 	fmt.Println("c:", c)

// 	a[2] = 99

// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// 	fmt.Println("c:", c)
// }

// func main() {
// 	a := []int{1, 2}
// 	b := append(a, 3)
// 	c := append(a, 4)

// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// 	fmt.Println("c:", c)
// }

// func main() {
// 	a := []int{10, 20, 30}
// 	b := a
// 	b[0] = 99

// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// }

// func main() {
// 	a := []int{1, 2, 3}
// 	b := make([]int, len(a))
// 	copy(b, a)
// 	b[0] = 100

// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// }

func main() {
	a := make([]int, 0, 2)
	a = append(a, 1)
	b := a
	a = append(a, 2)
	b = append(b, 3)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
