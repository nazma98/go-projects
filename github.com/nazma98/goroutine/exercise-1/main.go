package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Print(i)
		}()
	}
	time.Sleep(5 * time.Second)
}
