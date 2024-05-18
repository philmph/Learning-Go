package main

import "fmt"

func main() {
	// Constants (not written in CAPS in Go) - Same as var declaration
	// Unexported seemingly best practice in camelCase versus exportet
	// in PascalCase
	const constInt int = 123

	// Strongly typed
	var i int = 123

	var f float32 = 12.3
	var ff float64 = 12.3

	var s string = "My string"

	var b bool = false

	/*
		Defaults for %v from docs:
		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
	*/

	fmt.Printf("Explicitly defined (ddggst): %d, %d, %g, %g, %s, %t\n", constInt, i, f, ff, s, b)
	fmt.Printf("Is the same as     (vvvvvv): %v, %v, %v, %v, %v, %v\n", constInt, i, f, ff, s, b)

	// Implicit assignment
	i2 := 123
	ff2 := 12.3 //
	s2 := "My string"
	b2 := false

	fmt.Printf("Implcitly assigned: %v, %v, %v, %v\n", i2, ff2, s2, b2)
	fmt.Printf("Note that ff2 will default to float64: %T\n", ff2)
}
