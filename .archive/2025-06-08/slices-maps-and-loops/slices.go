package main

import "fmt"

func slicesFunc() {
	// Syntax without make
	users := []string{"John", "Paul", "George", "Ringo"}

	fmt.Printf("1: Lenght: %d, Capability: %d Content: %s\n", len(users), cap(users), users)

	// Slice with 0 initialzed values
	makeUsers := make([]string, 0, 4)

	// Results in length of 1 for the slice
	makeUsers = append(makeUsers, "John")

	fmt.Printf("2: Lenght: %d, Capability: %d Content: %s\n", len(makeUsers), cap(makeUsers), makeUsers)

	// Slice with 4 initialzed values and cap of 4 (inherited)
	makeUsers = make([]string, 4)

	// Same as
	makeUsers = make([]string, 4, 4)

	fmt.Printf("3: Lenght: %d, Capability: %d Content: %s\n", len(makeUsers), cap(makeUsers), makeUsers)

	// Results in length of 5 for the slice because 4 are already initialized with string default value ""
	makeUsers = append(makeUsers, "John")

	fmt.Printf("4: Lenght: %d, Capability: %d Content: %s\n", len(makeUsers), cap(makeUsers), makeUsers)
}
