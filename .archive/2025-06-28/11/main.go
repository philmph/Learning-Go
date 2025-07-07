package main

import "fmt"

func calculateSomething(x, y int) int {
	return x * y
}

func main() {
	x, y := 5, 5
	fmt.Println("Hi: ", calculateSomething(x, y))
}
