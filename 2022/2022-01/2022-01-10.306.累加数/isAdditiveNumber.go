package main

import (
	"strconv"
	"strings"
)

func main() {

}

func isAdditiveNumber(num string) bool {
	l := len(num)

	for i := 1; i <= l/2; i++ {
		if num[0] == '0' && i > 1 {
			return false
		}

		num1, _ := strconv.Atoi(num[:i])

		for j := 1; i <= l-i-j && j <= l-i-j; j++ {
			if num[i] == '0' && j > 1 {
				break
			}

			num2, _ := strconv.Atoi(num[i : i+j])
			if isValid(num1, num2, num[i+j:]) {
				return true
			}
		}
	}
	return false
}

func isValid(n1, n2 int, num string) bool {
	if len(num) == 0 {
		return true
	}

	n3 := strconv.Itoa(n1 + n2)

	if !strings.HasPrefix(num, n3) {
		return false
	}
	return isValid(n2, n1+n2, num[len(n3):])
}
