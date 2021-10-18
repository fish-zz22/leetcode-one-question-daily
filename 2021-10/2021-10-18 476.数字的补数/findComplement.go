package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findComplement(7))
}

func findComplement(num int) int {
	var builder strings.Builder
	numBinaryStr := fmt.Sprintf("%b", num)

	for _, v := range numBinaryStr {
		if string(v) == "0" {
			builder.WriteString("1")
		} else {
			builder.WriteString("0")
		}
	}

	sum, _ := strconv.ParseInt(builder.String(), 2, 64)
	return int(sum)
}
