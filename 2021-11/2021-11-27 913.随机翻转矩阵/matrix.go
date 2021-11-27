package main

import (
	"math/rand"
)

func main() {

}

type Solution struct {
	m, n, total int
	hash        map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m, n, m*n - 1, map[int]int{}}
}

func (s *Solution) Flip() []int {
	r := rand.Intn(s.total + 1)
	idx, ok := s.hash[r]
	if !ok {
		idx = r
	}
	if v, okay := s.hash[s.total]; okay {
		s.hash[r] = v
	} else {
		s.hash[r] = s.total
	}
	s.total--
	return []int{idx / s.n, idx % s.n}
}

func (s *Solution) Reset() {
	s.total = s.m*s.n - 1
	s.hash = map[int]int{}
}
