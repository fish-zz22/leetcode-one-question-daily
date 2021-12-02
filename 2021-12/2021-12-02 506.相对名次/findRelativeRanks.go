package main

import (
	"sort"
	"strconv"
)

func main() {
	findRelativeRanks([]int{10, 3, 8, 9, 4})
}

func findRelativeRanks(score []int) []string {
	newScore := make([]int, len(score))
	newRankMap := make(map[int]int, len(score))
	rankString := make([]string, 0, len(score))
	rankMap := map[int]string{
		1: "Gold Medal",
		2: "Silver Medal",
		3: "Bronze Medal",
	}

	copy(newScore, score)
	sort.Slice(newScore, func(i, j int) bool {
		return newScore[i] > newScore[j]
	})

	for i, v := range newScore {
		newRankMap[v] = i + 1
	}

	for _, v := range score {
		if newRankMap[v] <= 3 {
			rankString = append(rankString, rankMap[newRankMap[v]])
		} else {
			rankString = append(rankString, strconv.Itoa(newRankMap[v]))
		}
	}
	return rankString
}
