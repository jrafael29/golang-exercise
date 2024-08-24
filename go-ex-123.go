package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
// You might find strings.Fields helpful.
// ref: https://go.dev/tour/moretypes/23

// solution by José rRafael:

func WordCount(s string) map[string]int {
	values := make(map[string]int)

	for _, v := range strings.SplitAfter(s, " ") {
		v = strings.Trim(v, " ")
		_, ok := values[v]
		if !ok {
			values[v] = 1
		} else {
			values[v] = values[v] + 1
		}
	}

	return values
}

func main() {
	wc.Test(WordCount)
}
