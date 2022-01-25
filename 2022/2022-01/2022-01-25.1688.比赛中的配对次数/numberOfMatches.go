package main

import "fmt"

func main() {
	fmt.Println(7 / 2)
	fmt.Println(numberOfMatches(14))
}

func numberOfMatches(n int) (res int) {
	for n != 1 {
		res += n / 2

		if n%2 == 0 {
			n = n / 2
		} else {
			n = n/2 + 1
		}
	}
	return
}
