package main

import "fmt"


func EvenOrOdd(x int){
    if x % 2 == 1{
        fmt.Println("ODD")
    }else{
        fmt.Println("EVEN")
    }
}

func main(){
    EvenOrOdd(5)
}