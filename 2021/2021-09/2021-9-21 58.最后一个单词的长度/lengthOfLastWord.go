package main

func main() {
	s := "   fly me   to   the moon  "
	lengthOfLastWord(s)
}

func lengthOfLastWord(s string) int {
	l := len(s) - 1
	var res int
	for i := l; i >= 0; i-- {
		if string(s[i]) != " " {
			res++
		} else if res != 0 {
			return res
		}
	}
	return res
}
