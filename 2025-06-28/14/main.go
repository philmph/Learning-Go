package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	words := strings.Split(s, " ")

	out := make([]string, len(words))

	return strings.Join(out, " ")
}

func main() {
	s := "This is a sentence"
	fmt.Println(reverseString(s))
}
