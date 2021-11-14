package main

func main() {

}

func detectCapitalUse(word string) bool {
	count := 0
	for _, r := range word {
		if r <= 'Z' && r >= 'A' {
			count++
		}
	}
	return count == 0 || count == len(word) || (count == 1 && word[0] <= 90 && word[0] >= 65)
}
