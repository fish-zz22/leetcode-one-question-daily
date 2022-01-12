package main

import "math"

func main() {

}

func increasingTriplet(nums []int) bool {
	i, j := math.MaxInt64, math.MaxInt64

	for _, v := range nums {
		if v <= i {
			i = v
		} else if v <= j {
			j = v
		} else {
			return true
		}
	}
	return false
}
