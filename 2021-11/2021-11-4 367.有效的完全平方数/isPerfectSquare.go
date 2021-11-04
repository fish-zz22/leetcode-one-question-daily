package main

func main() {

}

func isPerfectSquare(num int) bool {
	for temp := 1; temp <= num; temp++ {
		if temp*temp > num {
			return false
		} else if temp*temp == num {
			return true
		}
	}
	return false
}
