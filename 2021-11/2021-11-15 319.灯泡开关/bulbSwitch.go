package main

import (
	"fmt"
	"math"
)

func main() {
	n := 4
	fmt.Println(bulbSwitch(n))
}

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
