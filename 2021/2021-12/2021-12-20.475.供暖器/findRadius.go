package main

import (
	"math"
	"sort"
)

func main() {

}

func findRadius(houses []int, heaters []int) (res int) {
	sort.Ints(heaters)
	for _, house := range houses {

		temp := math.MaxInt64
		index := sort.SearchInts(heaters, house+1)

		if index < len(heaters) {
			temp = heaters[index] - house
		}

		if index-1 >= 0 && house-heaters[index-1] < temp {
			temp = house - heaters[index-1]
		}

		if temp > res {
			res = temp
		}
	}
	return
}
