package main

import "fmt"

var (
	myString string = "myString"
)

func main() {
	users := make(map[string]int, 5)

	users["philipp"] = 33
	users["thomas"] = 3

	user := "thomas"

	age, ok := users[user]
	if !ok {
		fmt.Println(ok)
		fmt.Printf("User %v not found\n", user)
	} else {
		fmt.Println(ok)
		fmt.Printf("User %v is %d years old\n", user, age)
	}

	age, ok = users["noone"]
	if !ok {
		fmt.Println(ok)
		fmt.Printf("User noone not found\n")
	} else {
		fmt.Println(ok)
		fmt.Printf("User %v is %d years old\n", user, age)
	}

	for k, age := range users {
		fmt.Printf("User %v age: %v\n", k, age)
	}

	dogs := []string{"wadu", "hek"}

	for i, name := range dogs {
		fmt.Printf("Index %d name: %v\n", i, name)
	}

	slice := make([]int, 1)

	fmt.Println(len(slice), slice[0])

	fmt.Println(myString)
}
