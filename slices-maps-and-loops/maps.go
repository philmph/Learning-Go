package main

import "fmt"

func maps() {
	users := map[int]string{1: "John", 2: "Paul", 3: "George", 4: "Ringo"}

	fmt.Printf("1: Lenght: %d, Content: %+v\n", len(users), users)

	// Initialize a map with make function with automaitc memory allocation
	var customUsers map[int]string
	customUsers2 := map[int]string{}
	makeUsers := make(map[int]string)

	fmt.Printf("2.1: Lenght: %d, Content: %+v\n", len(customUsers), customUsers)
	fmt.Printf("2.2: Lenght: %d, Content: %+v\n", len(customUsers2), customUsers2)
	fmt.Printf("2.3: Lenght: %d, Content: %+v\n", len(makeUsers), makeUsers)

	// Same as make(map[int]string) and makes no sense to me
	makeUsers = make(map[int]string, 0)

	fmt.Printf("3: Lenght: %d, Content: %+v\n", len(makeUsers), makeUsers)

	// Also same but allocates more memory (?)
	// Also makes no sense to me
	makeUsers = make(map[int]string, 4)

	fmt.Printf("4: Lenght: %d, Content: %+v\n", len(makeUsers), makeUsers)

	// Assign values
	makeUsers[1] = "John"
	makeUsers[2] = "Paul"
	makeUsers[3] = "George"
	makeUsers[4] = "Ringo"

	// Added 4?
	fmt.Printf("5: Lenght: %d, Content: %+v\n", len(makeUsers), makeUsers)
}
