package main

import "fmt"

func main() {
	timeSeries := []int{1, 2, 7}
	duration := 3
	fmt.Println(findPoisonedDuration(timeSeries, duration))
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	var seconds int
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] >= duration {
			seconds += duration
		} else {
			seconds += timeSeries[i] - timeSeries[i-1]
		}
	}
	seconds += duration
	return seconds
}
