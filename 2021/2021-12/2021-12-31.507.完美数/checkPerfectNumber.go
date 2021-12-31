package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPerfectNumber(2))
}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	res := 1

	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			res += i
			res += num / i
		}
	}

	if res == num {
		return true
	}
	return false
}
