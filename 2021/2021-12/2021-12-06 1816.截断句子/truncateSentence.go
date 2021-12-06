package main

import (
	"fmt"
)

func main() {
	s := "Hello how are you Contestant"
	k := 4
	fmt.Println(truncateSentence(s, k))
}

func truncateSentence(s string, k int) string {
	//var builder strings.Builder
	//
	//strList := strings.Split(s, " ")
	//for i := 0; i < k; i++ {
	//	if i == k-1 {
	//		builder.WriteString(strList[i])
	//	} else {
	//		builder.WriteString(strList[i])
	//		builder.WriteString(" ")
	//	}
	//
	//}
	//return builder.String()

	index := -1
	for i, v := range s {
		if v == ' ' {
			k--
		}

		if k == 0 {
			index = i
			break
		}
	}
	if index == -1 {
		return s
	}
	return s[:index]
}
