package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	fmt.Println(uncommonFromSentences(s1, s2))
}

func uncommonFromSentences(s1 string, s2 string) (res []string) {
	var builder strings.Builder

	builder.WriteString(s1)
	builder.WriteString(" ")
	builder.WriteString(s2)

	l := strings.Split(builder.String(), " ")

	hash := make(map[string]int)
	for _, v := range l {
		hash[v]++
	}

	for k, v := range hash {
		if v == 1 {
			res = append(res, k)
		}
	}
	return
}
