package main

import "bytes"

func main() {

}

func originalDigits(s string) string {
	hash := make(map[rune]int, len(s))
	for _, ch := range s {
		hash[ch]++
	}

	count := [10]int{}
	count[0] = hash['z']
	count[2] = hash['w']
	count[4] = hash['u']
	count[6] = hash['x']
	count[8] = hash['g']

	count[3] = hash['h'] - count[8]
	count[5] = hash['f'] - count[4]
	count[7] = hash['s'] - count[6]

	count[1] = hash['o'] - count[0] - count[2] - count[4]

	count[9] = hash['i'] - count[5] - count[6] - count[8]
	ans := []byte{}
	for i, c := range count {
		ans = append(ans, bytes.Repeat([]byte{byte('0' + i)}, c)...)
	}
	return string(ans)
}
