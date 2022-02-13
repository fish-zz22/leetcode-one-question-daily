package main

func main() {

}

func countGoodRectangles(rectangles [][]int) (ans int) {
	maxLen := 0
	for _, rect := range rectangles {
		k := min(rect[0], rect[1])
		if k == maxLen {
			ans++
		} else if k > maxLen {
			maxLen, ans = k, 1
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
