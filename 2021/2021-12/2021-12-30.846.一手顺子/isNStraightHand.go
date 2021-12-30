package main

import "sort"

func main() {

}

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)

	if n%groupSize != 0 {
		return false
	}

	sort.Ints(hand)
	count := make(map[int]int)

	for _, num := range hand {
		count[num]++
	}

	for _, x := range hand {
		if count[x] == 0 {
			continue
		}

		for num := x; num < x+groupSize; num++ {
			if count[num] == 0 {
				return false
			}
			count[num]--
		}
	}
	return true
}
