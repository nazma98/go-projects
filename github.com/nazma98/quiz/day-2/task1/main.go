package main

import (
    "fmt"
    "time"
)

func Sum(x int) int {
    sum := 0
    for i := 1; i <= x; i++ {
        sum += i
    }
    return sum
}

func main() {
    x := 100

    start := time.Now()
    result := Sum(x)
    elapsed := time.Since(start)

    fmt.Println("Sum of first", x, "numbers:", result)
    fmt.Println("Execution Time:", elapsed)
}