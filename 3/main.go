package main

import "fmt"

func main() {
	users := make(map[string]int, 5)

	users["philipp"] = 33
	users["elsie"] = 3

	user := "elsie"

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

	dogs := []string{"philipp", "elsie"}

	for i, name := range dogs {
		fmt.Printf("Index %d name: %v\n", i, name)
	}

	slice := make([]int, 1)

	fmt.Println(len(slice), slice[0])
}
