package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	words := strings.Split(s, " ")
	ret := make([]string, 0, len(words))

	for _, w := range words {
		runes := []rune(w)

		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		ret = append(ret, string(runes))
	}

	return strings.Join(ret, " ")
}

func main() {
	s := "This is a sentence"
	fmt.Println(reverseString(s))
}
