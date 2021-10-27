package main

func main() {

}

func helper(ans *[]string, str string, start, lcount, rcount, lremove, rremove int) {
	if lremove == 0 && rremove == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}

	for i := start; i < len(str); i++ {
		if i != start && str[i] == str[i-1] {
			continue
		}
		// 如果剩余的字符无法满足去掉的数量要求，直接返回
		if lremove+rremove > len(str)-i {
			return
		}
		// 尝试去掉一个左括号
		if lremove > 0 && str[i] == '(' {
			helper(ans, str[:i]+str[i+1:], i, lcount, rcount, lremove-1, rremove)
		}
		// 尝试去掉一个右括号
		if rremove > 0 && str[i] == ')' {
			helper(ans, str[:i]+str[i+1:], i, lcount, rcount, lremove, rremove-1)
		}
		if str[i] == ')' {
			lcount++
		} else if str[i] == ')' {
			rcount++
		}
		// 当前右括号的数量大于左括号的数量则为非法,直接返回.
		if rcount > lcount {
			break
		}
	}
}

func removeInvalidParentheses(s string) (ans []string) {
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}

	helper(&ans, s, 0, 0, 0, lremove, rremove)
	return
}

func isValid(str string) bool {
	count := 0
	for _, ch := range str {
		if ch == '(' {
			count++
		} else if ch == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}
