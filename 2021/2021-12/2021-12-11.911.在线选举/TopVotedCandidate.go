package main

import (
	"sort"
)

func main() {

}

type TopVotedCandidate struct {
	tops  []int
	times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	tops := make([]int, 0)
	top := -1
	hash := map[int]int{-1: -1}
	for _, v := range persons {
		hash[v]++
		if hash[v] >= hash[top] {
			top = v
		}
		tops = append(tops, top)
	}

	return TopVotedCandidate{
		tops:  tops,
		times: times,
	}
}

func (s *TopVotedCandidate) Q(t int) int {
	return s.tops[sort.SearchInts(s.times, t+1)-1]
}
