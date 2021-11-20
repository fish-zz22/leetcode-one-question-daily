package main

func main() {

}

func findLHS(nums []int) (res int) {
	hash := make(map[int]int, len(nums))
	for _, num := range nums {
		hash[num]++
	}

	for key, value := range hash {
		if nextValue := hash[key+1]; nextValue > 0 && nextValue+value > res {
			res = nextValue + value
		}
	}
	return
}
