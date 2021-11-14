package main

import "strings"

func main() {

}

type MapSum struct {
	hash map[string]int
}

func Constructor() MapSum {
	return MapSum{
		hash: make(map[string]int, 0),
	}
}

func (m *MapSum) Insert(key string, val int) {
	m.hash[key] = val
}

func (m *MapSum) Sum(prefix string) int {
	sum := 0
	for key, val := range m.hash {
		if strings.HasPrefix(key, prefix) {
			sum += val
		}
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
