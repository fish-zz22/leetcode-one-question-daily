package main

import "fmt"

func main() {
	fmt.Println(20/7 + 20%7)
	fmt.Println(totalMoney(7))
}

func totalMoney(n int) (res int) {
	for i := 1; i <= n; i++ {
		if i%7 == 0 {
			res += i/7 - 1 + 7
		} else {
			res += i/7 + i%7
		}
	}
	return res
}
