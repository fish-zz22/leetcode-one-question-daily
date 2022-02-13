package main

import (
	"fmt"
	"math"
)

func main() {
	text := "leetcode"
	fmt.Println(maxNumberOfBalloons(text))
}

func maxNumberOfBalloons(text string) int {
	var a, b, l, o, n int

	for _, v := range text {
		switch v {
		case 'a':
			a++
		case 'b':
			b++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		}
	}

	l /= 2
	o /= 2

	res := math.MaxInt64
	if a < res {
		res = a
	}
	if b < res {
		res = b
	}
	if l < res {
		res = l
	}
	if o < res {
		res = o
	}
	if n < res {
		res = n
	}
	return res
}
