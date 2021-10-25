package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(matrix, 1))
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix)-1, len(matrix[0])-1
	if target < matrix[0][0] || target > matrix[m][n] {
		return false
	}

	for i, nums := range matrix {
		if target > matrix[i][n] {
			continue
		}

		for _, num := range nums {
			if num == target {
				return true
			}
		}
	}
	return false
}
