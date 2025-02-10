package main
import "fmt"


func add(a, b int){
    return a+b
}

func substract(a, b int){
    return a-b
}

func multiply(a, b int){
    return a*b
}


func main(){
    
    fmt.Println(add(5,5))
    fmt.Println(substract(5,5))
    fmt.Println(multiply(5,5))
    fmt.Println(add(5,5))

}