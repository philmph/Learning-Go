package main

import "fmt"

type User struct {
	health int
	name   string
	alive  bool
}

func main() {
	user := &User{
		100,
		"John",
		true,
	}

	var user2 *User = &User{
		50,
		"Bill",
		false,
	}

	fmt.Printf("User: %v\n", *user)
	fmt.Printf("User: %v\n", *user2)
}
