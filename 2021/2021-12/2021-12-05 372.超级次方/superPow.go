package main

func main() {

}

const mod = 1337

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 != 0 {
			res = res * x % mod
		}

		x = x * x % mod
	}
	return res
}

func superPow(a int, b []int) int {
	res := 1
	for i := len(b) - 1; i >= 0; i-- {
		res = res * pow(a, b[i]) % mod
		a = pow(a, 10)
	}
	return res
}
