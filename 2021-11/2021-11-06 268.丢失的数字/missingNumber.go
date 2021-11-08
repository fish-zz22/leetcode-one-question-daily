package main

func main() {

}

func missingNumber(nums []int) int {
	l := len(nums)
	sum := l * (l + 1) / 2
	arrSum := 0
	for _, v := range nums {
		arrSum += v
	}
	return sum - arrSum
}
