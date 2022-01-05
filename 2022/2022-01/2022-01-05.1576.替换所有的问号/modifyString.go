package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	s := "????????????????????????????????????????????????????????????????????????????????????????????????????"
	fmt.Println(modifyString(s), len(s))
}

func modifyString(s string) string {
	var builder strings.Builder
	var ch string

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i, v := range s {
		if v == '?' {
			ch = string('a' + uint8(r.Intn(24)))

			if len(s) == 1 {
				builder.WriteString(ch)
				break
			}

			if i == 0 {
				for ch == string(s[i+1]) {
					ch = string('a' + uint8(r.Intn(24)))
				}
				builder.WriteString(ch)
			} else if i == len(s)-1 {

				for ch == string(builder.String()[i-1]) || ch == string(s[i-1]) {
					ch = string('a' + uint8(r.Intn(24)))
				}

				builder.WriteString(ch)
			} else {
				if s[i-1] == '?' {
					for ch == string(s[i+1]) || ch == string(builder.String()[i-1]) {
						ch = string('a' + uint8(r.Intn(24)))
					}
				} else {
					for ch == string(s[i+1]) || ch == string(s[i-1]) {
						ch = string('a' + uint8(r.Intn(24)))
					}
				}

				builder.WriteString(ch)
			}

		} else {
			builder.WriteString(string(v))
		}
	}

	return builder.String()
}
