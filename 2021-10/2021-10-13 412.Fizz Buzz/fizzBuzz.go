package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	var strList []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			strList = append(strList, "FizzBuzz")
		} else if i%3 == 0 {
			strList = append(strList, "Fizz")
		} else if i%5 == 0 {
			strList = append(strList, "Buzz")
		} else {
			strList = append(strList, strconv.Itoa(i))
		}
	}
	return strList
}
