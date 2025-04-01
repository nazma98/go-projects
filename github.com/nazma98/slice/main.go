package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "Go", "Interview", "question"}
	fmt.Println(arr)

	s := arr[1:4]
	fmt.Println(s)
	fmt.Println(len(s), " ", cap(s))

	s1 := s[1:2]
	fmt.Println("Slice :",s1, " Len: ", len(s1), " cap: ", cap(s1))

	s2 := []string{"I", "am", "learning", "Go"}
	fmt.Println("Slice :",s2, " Len: ", len(s2), " cap: ", cap(s2))

	s3 := make([]int, 3, 5)
	s3[0] = 5
	s3[2] = 10
	fmt.Println(s3)
	fmt.Println("slice ", s3, "len ", len(s3), "cap ", cap(s3))
}