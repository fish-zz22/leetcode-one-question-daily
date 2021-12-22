package main

import "strings"

func main() {

}

func repeatedStringMatch(a string, b string) int {
	var builder strings.Builder
	res := 1
	builder.WriteString(a)

	for len(builder.String()) < len(b) {
		builder.WriteString(a)
		res++
	}

	if strings.Contains(builder.String(), b) {
		return res
	}

	builder.WriteString(a)
	res++

	if strings.Contains(builder.String(), b) {
		return res
	}
	return -1
}
