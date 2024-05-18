package main

import "fmt"

func main() {
	i := 123

	pointer := &i
	fmt.Printf("Pointer value and type: %v, %[1]T\n", pointer)

	// Change value using pointer instead of i directly
	*pointer = 321
	fmt.Println(i, *pointer)
}
