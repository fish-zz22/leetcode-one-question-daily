package main

func main() {

}

func kInversePairs(n int, k int) int {
	var mod int64 = 1000000007
	dp := make([][]int64, k+1)
	dp[0] = make([]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < len(dp); i++ {
		dp[i] = make([]int64, n+1)
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % mod
			if i >= j {
				dp[i][j] = (dp[i][j] + mod - dp[i-j][j-1]) % mod
			}
		}
	}
	return int(dp[k][n])
}
