package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	dictionary := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		_, isExist := dictionary[word]
		if isExist {
			dictionary[word]++ 
		} else {
			dictionary[word] = 1
		}
	}
	return dictionary
}

func main() {
	wc.Test(WordCount)
}
