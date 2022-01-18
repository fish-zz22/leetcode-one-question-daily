package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	timePoints := []string{"00:00", "23:59", "00:00"}
	fmt.Println(findMinDifference(timePoints))
}

func findMinDifference(timePoints []string) int {
	res := math.MaxInt64

	sort.Slice(timePoints, func(i, j int) bool {
		return timePoints[i] < timePoints[j]
	})

	t0Minutes := getMinutes(timePoints[0])
	preMinutes := t0Minutes
	for _, t := range timePoints[1:] {
		minutes := getMinutes(t)
		res = min(res, minutes-preMinutes)
		preMinutes = minutes
	}
	res = min(res, t0Minutes+1440-preMinutes)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinutes(s string) int {
	return (int(s[0]-'0')*10+int(s[1]-'0'))*60 + int(s[3]-'0')*10 + int(s[4]-'0')
}
