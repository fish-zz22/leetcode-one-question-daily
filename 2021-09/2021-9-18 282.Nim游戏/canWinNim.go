package main

import "fmt"

func main() {
	fmt.Println(canWinNim(8))
}

func canWinNim(n int) bool {
	if n%4 != 0 {
		return true
	}
	return false
}
