package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "alice is a good girl she is a good student"
	first := "a"
	second := "good"
	fmt.Println(findOcurrences(text, first, second))
}

func findOcurrences(text string, first string, second string) (res []string) {
	words := strings.Split(text, " ")
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			res = append(res, words[i])
		}
	}
	return
}
