package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	points := [][]int{
		{2, 1},
		{2, 2},
		{2, 3},
	}
	angle := 90
	location := []int{1, 1}
	fmt.Println(visiblePoints(points, angle, location))
}

func visiblePoints(points [][]int, angle int, location []int) int {
	count, maxCount := 0, 0
	degrees := make([]float64, 0)
	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			count++
		} else {
			degrees = append(degrees, math.Atan2(float64(p[0]-location[0]), float64(p[1]-location[1])))
		}
	}
	sort.Float64s(degrees)

	n := len(degrees)
	for _, deg := range degrees {
		degrees = append(degrees, deg+2*math.Pi)
	}
	degree := float64(angle) * math.Pi / 180

	for i, deg := range degrees[:n] {
		j := sort.Search(n*2, func(j int) bool {
			return degrees[j] > deg+degree
		})
		if j-i > maxCount {
			maxCount = j - i
		}
	}
	return maxCount + count
}
