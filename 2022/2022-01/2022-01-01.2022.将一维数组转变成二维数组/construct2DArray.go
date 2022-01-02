package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4}
	fmt.Println(construct2DArray(original, 2, 2))
	fmt.Println(2 % 1)
}

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}

	res := make([][]int, 0, m-1)

	for i := 0; i < len(original); i += n {
		res = append(res, original[i:i+n])
	}
	return res
}
