package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i * i)
	}

	for {
		fmt.Println("Hello")
		break
	}

	for i := 1; i <= 5; i++ {
		if i == 2 {
			continue
		} else if i == 3 {
			continue
		} else {
			fmt.Println(i * i)
		}
	}

	for i := 1; i <= 5; i++ {
		if i == 2 || i == 3 {
			continue
		} else {
			fmt.Println(i * i)
		}
	}
}
