package main

import "math/rand"

func main() {

}

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (s *Solution) Reset() []int {
	return s.nums
}

func (s *Solution) Shuffle() []int {
	newNums := make([]int, len(s.nums))
	copy(newNums, s.nums)
	rand.Shuffle(len(newNums), func(i, j int) {
		newNums[i], newNums[j] = newNums[j], newNums[i]
	})
	return newNums
}
