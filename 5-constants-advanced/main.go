package main

import "fmt"

const c1 int = 1
const c2 = 2

// Not possible for constants
// const c3 := 3

// Also not possible
// const c4 int
// c4 = 4

func main() {
	const c5 int = 5

	fmt.Println(c1, c2, c5)
}
