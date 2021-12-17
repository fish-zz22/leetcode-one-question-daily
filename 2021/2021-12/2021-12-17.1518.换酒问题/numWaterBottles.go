package main

import "fmt"

func main() {
	fmt.Println(numWaterBottles(100, 3))
}

func numWaterBottles(numBottles int, numExchange int) (sum int) {
	sum = numBottles + numBottles/numExchange
	for numBottles >= numExchange {
		sum += (numBottles%numExchange + numBottles/numExchange) / numExchange
		numBottles = numBottles%numExchange + numBottles/numExchange
	}
	return
}
