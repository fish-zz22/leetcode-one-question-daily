package main

func main() {

}

func isPerfectSquare(num int) bool {
	temp := 1
	for temp <= num {
		if temp*temp > num {
			return false
		} else if temp*temp == num {
			return true
		} else {
			temp++
		}
	}
	return false
}
