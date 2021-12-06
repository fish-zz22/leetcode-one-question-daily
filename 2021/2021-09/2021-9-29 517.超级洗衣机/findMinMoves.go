package main

import (
	"math"
)

func main() {
	machines := []int{1, 0, 5}
	findMinMoves(machines)
}

func findMinMoves(machines []int) int {
	var temp, res int

	length, sum := machinesSum(machines)
	if sum%length != 0 {
		return -1
	}

	avg := sum / length
	for _, value := range machines {
		temp += value - avg
		// 从左右的洗衣机中取衣服和当前洗衣机需要转移的衣服取最大值
		res = max(res, max(int(math.Abs(float64(temp))), int(math.Abs(float64(value-avg)))))
	}
	return res
}

func machinesSum(machines []int) (int, int) {
	length, sum := 0, 0
	for _, value := range machines {
		length++
		sum += value
	}
	return length, sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
