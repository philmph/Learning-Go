package main

import "fmt"

type User struct {
	health int
	name   string
	alive  bool
}

func main() {
	slice := make([]int, 5, 10)
	sliceAlternative := []int{1, 2, 3, 4, 5}
	array := [5]int{1, 2, 3, 4, 5}

	slice = append(slice, 6)
	sliceAlternative = append(sliceAlternative, 6)

	newArray := append(array[:], 6)

	fmt.Printf("%+v\n", slice)
	fmt.Printf("%+v\n", sliceAlternative)
	fmt.Printf("%+v\n", array)
	fmt.Printf("%+v\n", newArray)
}
