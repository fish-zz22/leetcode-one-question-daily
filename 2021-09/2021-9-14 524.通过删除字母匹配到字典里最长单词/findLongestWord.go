package main

import (
	"fmt"
)

func main() {
	s := "bab"
	dictionary := []string{
		"ba", "ab", "a", "b",
	}
	res := findLongestWord(s, dictionary)
	fmt.Println(res)

}

func findLongestWord(s string, dictionary []string) string {
	//var strList []string
	//for _,str := range dictionary{
	//	if TwoPointsEqualsStrings(s,str){
	//		strList = append(strList,str)
	//	}
	//}
	//sort.Slice(strList, func(i, j int) bool {
	//	if len(strList[i])==len(strList[j]){
	//		return strList[i]<strList[j]
	//	}
	//	return len(strList[i])>len(strList[j])
	//})
	//if len(strList) != 0{
	//	return strList[0]
	//}
	//return ""

	//优化  使用中间变量
	var res string
	for _, str := range dictionary {
		if TwoPointsEqualsStrings(s, str) {
			if len(res) < len(str) || len(res) == len(str) && res > str {
				res = str
			}
		}
	}
	return res
}

func TwoPointsEqualsStrings(s, dictionaryStr string) bool {
	var i, j int
	if len(s) < len(dictionaryStr) {
		return false
	}

	for _, str := range s {
		// 双指针判断s字符串删除某些字符组成dictionaryStr
		fmt.Println(string(dictionaryStr[j]), string(str))
		if string(dictionaryStr[j]) == string(str) {
			i++
			j++
		} else {
			i++
		}

		if i > len(s) {
			return false
		} else if j == len(dictionaryStr) {
			return true
		}
	}
	return false
}
