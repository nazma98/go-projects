package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[row]) - 1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}

func main() {
	mat := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {3, 30, 34, 60}, {4, 32, 40, 65}}

	fmt.Println(searchMatrix(mat, 41))
}
