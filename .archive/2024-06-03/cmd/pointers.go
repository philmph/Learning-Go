package main

import "fmt"

func pointers() {
	// Define a variable
	s := "text"

	// Define e pointer
	var p *string

	// Assign p the memory address of s. Use & to get the address of a non pointer variable
	p = &s

	// Outputs
	fmt.Println("Normal variable value access:", s)
	fmt.Println("Normal variable memory address using &:", &s)
	fmt.Println("Pointer variable memory address (pointer is a memory address by default):", p)
	fmt.Println("Pointer variable value access using *:", *p)
}
