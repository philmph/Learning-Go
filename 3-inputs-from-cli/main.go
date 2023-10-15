package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)

	fmt.Println("Enter name and age:")
	count, err := fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Count: %v, Err: %v, Name: %v, Age: %v\n", count, err, name, age)

	fmt.Printf("Error is an %T", err)
}
