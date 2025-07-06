package main

import (
	"fmt"
	"strings"
)

type transformFunc func(string) string

func transformString(s string, fn transformFunc) string {
	return fn(s)
}

func transformToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	fmt.Println(transformString("Hello World!", transformToUpper))
}
