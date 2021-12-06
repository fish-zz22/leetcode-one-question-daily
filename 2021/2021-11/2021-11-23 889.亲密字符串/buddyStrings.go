package main

func main() {
	s := "ab"
	goal := "ba"
	buddyStrings(s, goal)
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	var char [26]int
	first, end := -1, -1
	for index, value := range s {
		char[value-'a']++
		if s[index] != goal[index] {
			if first == -1 {
				first = index
			} else if end == -1 {
				end = index
			} else {
				return false
			}
		}
	}

	if first == -1 && end == -1 {
		for _, v := range char {
			if v > 1 {
				return true
			}
		}
	} else if first != -1 && end != -1 {
		if s[first] == goal[end] && s[end] == goal[first] {
			return true
		}
	}
	return false
}
