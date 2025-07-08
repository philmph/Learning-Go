package main

import "fmt"

func main() {
	resultch := make(chan string)

	go func() {
		result := <-resultch
		fmt.Println(result)
	}()

	resultch <- "Hi"
}
