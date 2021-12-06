package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins(8))
}

func arrangeCoins(n int) int {
	temp := n
	for i := 1; i <= n; i++ {
		temp = temp - i

		if temp < 0 {
			return i - 1
		}

		if temp == 0 {
			return i
		}
	}
	return 0
}
