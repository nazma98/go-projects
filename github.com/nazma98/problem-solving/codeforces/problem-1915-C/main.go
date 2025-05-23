package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInt() int {
	line := readLine()
	num, _ := strconv.Atoi(line)
	return num
}

func readInts(n int) []int {
	nums := []int{}
	for len(nums) < n {
		line := readLine()
		parts := strings.Fields(line)
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			nums = append(nums, num)
			if len(nums) == n {
				break
			}
		}
	}
	return nums
}

func canSquare(nums []int) string {
	sum := 0
	for _, val := range nums {
		sum = sum + val
	}
	sqr := math.Sqrt(float64(sum))
	sqrtInt := int(sqr)
	if sqrtInt*sqrtInt == sum {
		return "yes"
	}
	return "no"
}

func main() {
	defer writer.Flush()

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		arr := readInts(n)
		res := canSquare(arr)
		fmt.Fprintf(writer, "%s\n", res)
	}
}
