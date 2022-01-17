package main

import "fmt"

func main() {
	fmt.Println(countVowelPermutation(2))
}

func countVowelPermutation(n int) int {
	const num = 1e9 + 7
	a, e, i, o, u := 1, 1, 1, 1, 1
	for k := 2; k <= n; k++ {
		aEnd := (e + i + u) % num
		eEnd := (a + i) % num
		iEnd := (e + o) % num
		oEnd := (i) % num
		uEnd := (i + o) % num
		a = aEnd
		e = eEnd
		i = iEnd
		o = oEnd
		u = uEnd
	}
	return (a + e + i + o + u) % num
}
