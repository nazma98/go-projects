package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	res := make([]bool, len(candies))

	for _, val := range candies {
		if max < val {
			max = val
		}
	}

	for i, val := range candies {
		if val+extraCandies >= max {
			res[i] = true
		}
	}
	return res
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
