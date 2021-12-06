package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(-2147483648, 1))
}

func divide(dividend int, divisor int) int {
	var flag bool
	var res, count, base int

	if dividend == -2147483648 && divisor == -1 {
		return math.MaxInt32
	}

	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		flag = true
	} else {
		flag = false
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res = divisor
	for res <= dividend {
		base = 1
		count += base

		// 快速加法
		for res+res <= dividend {
			res += res
			count += base
			base += base
		}
		dividend -= res
		res = divisor
	}

	if !flag {
		count = -count
	}
	return count
}
