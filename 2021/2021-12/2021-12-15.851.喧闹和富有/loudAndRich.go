package main

func main() {

}

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	g := make([][]int, n)
	for _, r := range richer {
		g[r[1]] = append(g[r[1]], r[0])
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var dfs func(int)
	dfs = func(x int) {
		if ans[x] != -1 {
			return
		}
		ans[x] = x
		for _, y := range g[x] {
			dfs(y)
			if quiet[ans[y]] < quiet[ans[x]] {
				ans[x] = ans[y]
			}
		}
	}
	for i := 0; i < n; i++ {
		dfs(i)
	}
	return ans
}
