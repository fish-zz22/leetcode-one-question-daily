package main

func main() {

}

func kInversePairs(n int, k int) int {
	const mod = 1e9 + 7
	dp := make([][]int, k+1)
	dp[0] = make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			if i >= j {
				dp[i][j] = dp[i][j] - dp[i-1][j-1]
			}
			dp[i][j] = (dp[i][j] + mod) % mod
		}
	}
	return dp[n][k]
}
