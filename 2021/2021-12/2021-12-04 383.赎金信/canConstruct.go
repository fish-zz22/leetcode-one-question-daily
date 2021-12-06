package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	//hash := make(map[rune]int, len(magazine))
	//for _, v := range magazine {
	//	hash[v]++
	//}
	//
	//for _, v := range ransomNote {
	//	if hash[v] <= 0 {
	//		return false
	//	}
	//	hash[v]--
	//}
	//return true

	// 减少占用内存
	var char [26]int
	for _, v := range magazine {
		char[v-'a']++
	}

	for _, v := range ransomNote {
		if char[v-'a'] <= 0 {
			return false
		}
		char[v-'a']--
	}
	return true
}
