package main

import (
	"fmt"
)

func immutable_string() {
	// Original string
	original := "Hello, World!"

	// The following code is not possible:
	// original[12] = "?"

	fmt.Println("String:", original)

	// However, the variable can be re-assigned a new value
	// The old value will be inaccessible if not saved in another
	// string variable and therefor garbage collected.
	duplicate := original
	original = "Hello, Jupiter?"

	fmt.Println("'Saved' string:", duplicate)
	fmt.Println("Re-assigned string:", original)
}
