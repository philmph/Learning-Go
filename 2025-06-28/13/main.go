package main

import (
	"fmt"
	"strconv"
)

func main() {
	msgch := make(chan string, 128)

	for i := 0; i < 10; i++ {
		msgch <- ("Hello " + strconv.Itoa(i))
	}

	close(msgch)
	for msg := range msgch {
		fmt.Println(msg)
	}

	var outside string

	for {
		outside = "Hello"
		// inside := "Hello"
		break
	}

	fmt.Println(outside)

	// fmt.Println(inside)
}
