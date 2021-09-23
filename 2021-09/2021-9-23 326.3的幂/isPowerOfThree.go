package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(int(math.Floor(float64(2) / float64(3))))
	fmt.Println(2 / 3)
	fmt.Println(isPowerOfThree(27))
}

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
