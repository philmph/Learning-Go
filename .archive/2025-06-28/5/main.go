package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5}

	for n := range n[:3] {

		fmt.Println(n)
	}

}
