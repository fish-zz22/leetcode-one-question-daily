package main

import "strings"

func main() {

}

func maxProduct(words []string) (res int) {
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if !strings.ContainsAny(words[i], words[j]) && res < len(words[i])*len(words[j]) {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return
}
