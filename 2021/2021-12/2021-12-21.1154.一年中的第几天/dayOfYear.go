package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	date := "2004-03-01"
	fmt.Println(dayOfYear(date))
}

func dayOfYear(date string) (res int) {
	monthList := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, m, d := strings.Split(date, "-")[0], strings.Split(date, "-")[1], strings.Split(date, "-")[2]

	year, _ := strconv.Atoi(y)
	month, _ := strconv.Atoi(m)
	day, _ := strconv.Atoi(d)

	for i := 1; i < month; i++ {
		res += monthList[i-1]
	}
	res += day

	if !IsLeapYear(year) && month > 2 {
		res -= 1
	}

	return
}

func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}
