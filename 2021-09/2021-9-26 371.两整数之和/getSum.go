package main

func main() {
	getSum(1, 3)
}

func getSum(a int, b int) int {
	for b != 0 {
		temp := uint(a&b) << 1
		a ^= b
		b = int(temp)
	}
	return a
}
