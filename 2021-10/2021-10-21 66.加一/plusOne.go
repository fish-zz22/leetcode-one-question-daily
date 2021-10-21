package main

import "fmt"

func main() {
	digits := []int{9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	var flag bool
	l := len(digits) - 1
	for i := l; i >= 0; i-- {
		if flag == true || i == l {
			if digits[i]+1 < 10 {
				digits[i] = digits[i] + 1
				flag = false
			} else {
				digits[i] = 0
				flag = true
			}
		}
	}
	if flag == true {
		digits = append([]int{1}, digits...)
	}
	return digits
}
