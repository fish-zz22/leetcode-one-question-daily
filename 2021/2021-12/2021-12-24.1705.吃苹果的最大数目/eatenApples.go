package main

import "container/heap"

func main() {

}

func eatenApples(apples []int, days []int) (res int) {
	h := hp{}
	day := 0
	for ; day < len(apples); day++ {

		// 扔掉过期苹果
		for len(h) > 0 && h[0].endDay <= day {
			heap.Pop(&h)
		}

		if apples[day] > 0 {
			heap.Push(&h, pair{apple: apples[day], endDay: day + days[day]})
		}

		if len(h) > 0 {
			h[0].apple--
			if h[0].apple == 0 {
				heap.Pop(&h)
			}
			res++
		}
	}

	for len(h) > 0 {
		// 扔掉过期苹果
		for len(h) > 0 && h[0].endDay <= day {
			heap.Pop(&h)
		}

		if len(h) == 0 {
			break
		}

		p := heap.Pop(&h).(pair)
		num := min(p.endDay-day, p.apple)
		res += num
		day += num
	}
	return
}

type pair struct {
	endDay int
	apple  int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].endDay < h[j].endDay
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
