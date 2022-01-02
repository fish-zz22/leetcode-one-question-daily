package main

func main() {

}

func lastRemaining(n int) int {
	flag := true
	step, res := 1, 1
	for n > 1 {
		if flag || n%2 != 0 {
			res += step
		}
		flag = !flag
		step *= 2
		n /= 2
	}
	return res
}
