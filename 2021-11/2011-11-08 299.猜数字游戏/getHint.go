package main

import (
	"fmt"
	"strings"
)

func main() {
	secret := "1"
	guess := "1"
	fmt.Println(getHint(secret, guess))
}

func getHint(secret string, guess string) string {
	var (
		builder    strings.Builder
		bullsCount int
		cowsCount  int
	)
	secretMap := make(map[uint8]int, len(secret))
	guessMap := make(map[uint8]int, len(guess))

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bullsCount++
		} else {
			secretMap[secret[i]]++
			guessMap[guess[i]]++
		}
	}
	builder.WriteString(fmt.Sprintf("%vA", bullsCount))
	for v := range secretMap {
		cowsCount += min(secretMap[v], guessMap[v])
	}
	builder.WriteString(fmt.Sprintf("%vB", cowsCount))
	return builder.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
