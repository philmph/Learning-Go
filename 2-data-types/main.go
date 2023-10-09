package main

import "fmt"

func main() {
	// Constants (not written in CAPS in Go) - Same as var declaration
	const i int = 123

	// Strongly typed
	var ii int = 123
	var f float32 = 12.3
	var ff float64 = 12.3

	var s string = "My string"

	var b bool = false

	fmt.Printf("%d, %d, %g, %g, %s, %t", i, ii, f, ff, s, b)
}
