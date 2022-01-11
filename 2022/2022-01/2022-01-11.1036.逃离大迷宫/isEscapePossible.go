package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	block := [][]int{{0, 3}, {1, 0}, {1, 1}, {1, 2}, {1, 3}}
	source := []int{0, 0}
	target := []int{0, 2}
	fmt.Println(isEscapePossible(block, source, target))
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	if len(blocked) == 0 {
		return true
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(source, target, current []int, blockMap, visited map[string]bool) bool
	dfs = func(source, target, current []int, blockMap, visited map[string]bool) bool {
		if current[0] == target[0] && current[1] == target[1] {
			return true
		}

		if int(math.Abs(float64(source[0]-current[0])))+int(math.Abs(float64(source[1]-current[1]))) > 200 {
			return true
		}

		visited[strconv.Itoa(current[0])+"-"+strconv.Itoa(current[1])] = true

		for _, dir := range dirs {
			x, y := current[0]+dir[0], current[1]+dir[1]
			currentStr := strconv.Itoa(x) + "-" + strconv.Itoa(y)

			if x >= 0 && x < 1e6 && y >= 0 && y < 1e6 && !blockMap[currentStr] && !visited[currentStr] {
				if dfs(source, target, []int{x, y}, blockMap, visited) {
					return true
				}
			}
		}
		return false
	}

	blockedMap := make(map[string]bool)

	for _, block := range blocked {
		blockedMap[strconv.Itoa(block[0])+"-"+strconv.Itoa(block[1])] = true
	}

	return dfs(source, target, source, blockedMap, map[string]bool{}) && dfs(target, source, target, blockedMap, map[string]bool{})
}
