package main

import (
	"time"
)

func main() {
	day, month, year := 31, 8, 2019
	dayOfTheWeek(day, month, year)
}

func dayOfTheWeek(day int, month int, year int) string {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
}
