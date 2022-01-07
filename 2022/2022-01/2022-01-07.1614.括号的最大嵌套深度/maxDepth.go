package main

import "fmt"

func main() {
	s := "1"
	fmt.Println(maxDepth(s))
}

func maxDepth(s string) int {
	var (
		res, temp int
		stack     []int32
	)

	for _, v := range s {
		if v == '(' {
			stack = append(stack, v)
		} else if v == ')' {
			temp = len(stack)
			stack = stack[:len(stack)-1]
		}

		if temp > res {
			res = temp
		}
	}
	return res
}
