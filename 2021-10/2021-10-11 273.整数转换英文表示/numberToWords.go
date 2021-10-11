package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numberToWords(1234567891))
}

var (
	individual = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teen       = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	ty         = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	billion    = 1000000000
	million    = 1000000
	thousand   = 1000
)

func numberToWords(num int) string {
	var builder strings.Builder
	if num == 0 {
		return "Zero"
	}

	if num >= billion {
		b := num / billion
		num = num % billion
		builder.WriteString(build(b) + " Billion ")
	}

	if num >= million {
		m := num / million
		num = num % million
		builder.WriteString(build(m) + " Million ")
	}

	if num >= thousand {
		t := num / thousand
		num = num % thousand
		builder.WriteString(build(t) + " Thousand ")
	}

	builder.WriteString(build(num))

	return strings.TrimRight(builder.String(), " ")
}

func build(num int) string {
	var builder strings.Builder
	if num >= 100 {
		hundred := num / 100
		num = num % 100
		builder.WriteString(individual[hundred] + " " + "Hundred ")
	}

	if num >= 20 {
		ten := num / 10
		num = num % 10
		builder.WriteString(ty[ten] + " " + individual[num])
	} else if num >= 10 {
		builder.WriteString(teen[num%10])
	} else {
		builder.WriteString(individual[num])
	}

	return strings.TrimRight(builder.String(), " ")
}
