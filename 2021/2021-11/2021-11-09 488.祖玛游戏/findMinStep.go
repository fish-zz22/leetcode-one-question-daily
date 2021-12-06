package main

import "fmt"

func main() {

}

func findMinStep(board string, hand string) int {
	// BFS 暴力模拟 + 去重剪枝
	q := make([][]string, 0)
	q = append(q, []string{board, hand})
	depth := 0
	m := map[string]bool{} //用于去重剪枝
	for len(q) > 0 {
		depth++
		k := len(q)
		for k > 0 {
			k--
			cur := q[0]
			q = q[1:]
			// 把 h 中的每个元素暴力插入 b 中的每个位置，然后递归碰撞
			b, h := cur[0], cur[1]
			for i := 0; i < len(b); i++ {
				for j := 0; j < len(h); j++ {
					b2 := del3(b[0:i] + string(h[j]) + b[i:])
					h2 := h[0:j] + h[j+1:]
					if b2 == "" { // 递归碰撞后所有小球消失
						return depth
					}
					// 去重剪枝
					key := fmt.Sprintf("%v_%v", b2, h2)
					if m[key] {
						continue
					}
					m[key] = true
					q = append(q, []string{b2, h2})
				}
			}

		}
	}
	return -1
}

// 递归碰撞删除所有长度3及以上的子串
func del3(s string) string {
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			if count >= 3 {
				return del3(s[0:i-count] + s[i:])
			}
			count = 1
		}
	}
	if count >= 3 {
		return del3(s[0 : len(s)-count])
	}
	return s
}
