package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "xyxzxe"
	ch := byte('z')
	fmt.Println(reversePrefix(word, ch))
}

func reversePrefix(word string, ch byte) string {
	var index int

	for i, v := range word {
		if byte(v) == ch {
			index = i
			break
		}
	}

	if index == 0 {
		return word
	}

	var builder strings.Builder
	for i := index; i >= 0; i-- {
		builder.WriteString(string(word[i]))
	}
	builder.WriteString(word[index+1:])

	return builder.String()
}
