package main

func main() {

}

func grayCode(n int) []int {
	res := make([]int, 0)

	for i := 0; i < 1<<n; i++ {
		res = append(res, i>>1^i)
	}
	return res
}
