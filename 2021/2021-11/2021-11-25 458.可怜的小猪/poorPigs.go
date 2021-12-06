package main

import (
	"math"
)

func main() {

}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	pigs := 0
	for int(math.Pow(float64(minutesToTest/minutesToDie+1), float64(pigs))) < buckets {
		pigs += 1
	}
	return pigs
}
