package main

import "fmt"

// This file just showcases how go handles strings behind the scenes. Important to note is the special a at the end.
// https://www.youtube.com/watch?v=nxWqANttAdA

func main() {
	var t string = "testá"
	fmt.Printf("%T %[1]v\n", t)
	fmt.Printf("%T %[1]v\n", []rune(t))
	fmt.Printf("%T %[1]v\n", []byte(t))
}
