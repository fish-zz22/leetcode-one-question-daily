package main

import "fmt"

func main() {
	s := "tourist"
	fmt.Println(maxPower(s))
}

func maxPower(s string) int {
	res, temp := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			temp = 1
		} else {
			temp++
			res = max(temp, res)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
