package main

import "sort"

func main() {

}

type fraction struct {
	x, y int
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	l := len(arr)
	fractions := make([]fraction, 0, l*(l-1)/2)

	for i, x := range arr {
		for _, y := range arr[i+1:] {
			fractions = append(fractions, fraction{x: x, y: y})
		}
	}

	sort.Slice(fractions, func(i, j int) bool {
		return fractions[i].x*fractions[j].y < fractions[i].y*fractions[j].x
	})
	return []int{fractions[k-1].x, fractions[k-1].y}
}
