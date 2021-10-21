package main

import "fmt"

func main() {
	s := "AAAAAAAAAAA"
	findRepeatedDnaSequences(s)
}

func findRepeatedDnaSequences(s string) []string {
	var strList []string
	hash := make(map[string]int, 0)
	l := len(s)

	if l < 10 {
		return strList
	}

	for i := range s {
		j := i + 10
		if j > l {
			break
		}
		str := s[i:j]

		hash[str]++
		if hash[str] == 2 {
			strList = append(strList, str)
		}
		fmt.Println(strList, hash)
	}
	return strList
}
