package main

func main() {

}

type p struct {
	x, y int
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	var dirs = []p{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	preColor := grid[row][col]
	points := make([]p, 0, m*n)

	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		isBorder := false
		visit[x][y] = true

		for _, dir := range dirs {
			newX, newY := x+dir.x, y+dir.y
			if !(newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == preColor) {
				isBorder = true
			} else if !visit[newX][newY] {
				visit[newX][newY] = true
				dfs(newX, newY)
			}
		}
		if isBorder {
			points = append(points, p{x: x, y: y})
		}
	}
	dfs(row, col)

	for _, point := range points {
		grid[point.x][point.y] = color
	}
	return grid
}
