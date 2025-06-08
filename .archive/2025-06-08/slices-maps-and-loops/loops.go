package main

import "fmt"

func loopsFunc() {
	// Slices
	mySlice := []string{"John", "Paul", "George", "Ringo"}

	fmt.Println("Slices:")

	for iterator, object := range mySlice {
		fmt.Printf("Index: %d, Value: %s\n", iterator, object)
	}

	// Omitting one var
	for iterator := range mySlice {
		fmt.Printf("Value: %d\n", iterator)
	}

	// Omitting the iterator
	for _, object := range mySlice {
		fmt.Printf("Value: %s\n", object)
	}

	// Maps
	myMap := map[int]string{1: "John", 2: "Paul", 3: "George", 4: "Ringo"}

	fmt.Println("Maps:")

	for key, value := range myMap {
		fmt.Printf("Key: %d, Value: %s\n", key, value)
	}

	// Omitting one var
	for key := range myMap {
		fmt.Printf("Key: %d, \n", key)
	}

	// Omitting the key
	for _, value := range myMap {
		fmt.Printf("Key: %s, \n", value)
	}

	fmt.Println("Testing for key existance in map:")

	// Without check, accessinf a non existant key will return the default value of the type
	fmt.Printf("Key 0 (non existant): '%+v', Type: '%T'\n", myMap[0], myMap[0])

	// With check
	if value, ok := myMap[0]; ok {
		fmt.Printf("Key 0 (exist): '%+v', Type: '%T'\n", value, value)
	} else {
		fmt.Println("Key 0 does not exist")
	}

	// Can't access value and ok here outside if scope

	// Note that value and ok only exist in the if block
	if value, ok := myMap[1]; ok {
		fmt.Printf("Key 1 (exist): '%+v', Type: '%T'\n", value, value)
	} else {
		fmt.Println("Key 0 does not exist")
	}

	// Value and ok will exist outside the if block
	value, ok := myMap[1]
	if ok {
		fmt.Printf("Key 1 (exist): '%+v', Type: '%T'\n", value, value)
	} else {
		fmt.Println("Key 0 does not exist")
	}

	fmt.Printf("Variables still accessible: 'value': '%+v', Type: '%T', 'ok': '%+v', Type: '%T'\n", value, value, ok, ok)
}
