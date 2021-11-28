package main

func main() {
	s := "cbaebabacd"
	p := "abc"
	findAnagrams(s, p)
}

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return nil
	}
	var sList, pList [26]int
	var res []int

	for i, v := range p {
		pList[v-'a']++
		sList[s[i]-'a']++
	}
	if sList == pList {
		res = append(res, 0)
	}

	for i, v := range s[:sLen-pLen] {
		sList[v-'a']--
		sList[s[i+pLen]-'a']++
		if sList == pList {
			res = append(res, i+1)
		}
	}
	return res
}
