package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/../"
	fmt.Println(simplifyPath(s))
}

func simplifyPath(path string) string {
	stack := make([]string, 0)

	p := strings.Split(path, "/")
	for _, v := range p {

		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "" && v != "." {
			stack = append(stack, v)
		}
	}

	return "/" + strings.Join(stack, "/")
}
