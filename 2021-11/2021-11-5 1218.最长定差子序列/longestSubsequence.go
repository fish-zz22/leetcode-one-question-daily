package main

func main() {

}

func longestSubsequence(arr []int, difference int) int {
	var temp int
	hash := make(map[int]int, 0)
	for _, v := range arr {
		hash[v] = hash[v-difference] + 1
		temp = max(temp, hash[v])
	}
	return temp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
