package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u *user) greet() {
	fmt.Printf("%s says hello with %d years.\n", u.name, u.age)
}
