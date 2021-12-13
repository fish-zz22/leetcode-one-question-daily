package main

import "fmt"

func main() {
	grid := [][]int{
		//{3, 0, 8, 4},
		//{2, 4, 5, 7},
		//{9, 2, 6, 3},
		//{0, 3, 1, 0},
		{59, 88, 44},
		{3, 18, 38},
		{21, 26, 51},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

func maxIncreaseKeepingSkyline(grid [][]int) (sum int) {
	// 获取grid[i][j]中第i行第j列中imax和jmax
	// 然后遍历，取imax和jmax中两者较小的值 - grid[i][j]
	var iMax, jMax int
	m, n := len(grid), len(grid[0])
	skylines := make([]int, 0)

	for _, g := range grid {
		iMax = 0
		for _, v := range g {
			if v > iMax {
				iMax = v
			}
		}
		skylines = append(skylines, iMax)
	}

	for i := 0; i < n; i++ {
		jMax = 0
		for j := 0; j < m; j++ {
			if grid[j][i] > jMax {
				jMax = grid[j][i]
			}
		}
		skylines = append(skylines, jMax)
	}

	fmt.Println(skylines)
	for i, g := range grid {
		for j, v := range g {
			sum += min(skylines[i], skylines[m+j]) - v
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
