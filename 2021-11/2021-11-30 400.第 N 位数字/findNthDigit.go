package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(findNthDigit(11))
}

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	var (
		sum   = 0
		count = 9
		i     = 1
	)

	for ; sum+count*i < n; i++ {
		sum += count * i
		count *= 10
	}

	num := int(math.Pow10(i-1)) + (n-sum-1)/i
	index := (n - sum - 1) % i
	numStr := strconv.Itoa(num)
	res, _ := strconv.Atoi(string(numStr[index]))
	return res
}
