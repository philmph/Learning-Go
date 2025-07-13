package main

import (
	"fmt"
	"strings"
)

func reverseSentence(s string, sep string) string {
	words := strings.Split(s, sep)

	ret := make([]string, 0, len(words))

	for _, w := range words {
		ret = append(ret, string(reverseString(w)))
	}

	return strings.Join(ret, " ")
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	s := "This is a sentence"
	fmt.Println(reverseSentence(s, " "))
}
