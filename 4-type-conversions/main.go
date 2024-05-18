package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		i  int     = 123
		f  float64 = 12.3
		s  string  = "321"
		s2 string  = "321cba"
	)

	fmt.Printf("int to float: %v, %[1]T\n", float64(i))
	fmt.Printf("float to int: %v, %[1]T\n", int(f))

	iota := strconv.Itoa(i)
	fmt.Printf("int to string: %v, %[1]T\n", iota)

	atoi, err := strconv.Atoi(s)
	if err == nil {
		fmt.Printf("string to int: %v, %[1]T\n", atoi)
	} else {
		fmt.Printf("string to int: %v, %[1]T. Erroroutput: %v\n", atoi, err)
	}

	atoi2, err2 := strconv.Atoi(s2)
	if err2 == nil {
		fmt.Printf("string to int: %v, %[1]T\n", atoi2)
	} else {
		fmt.Printf("string to int: %v, %[1]T. Erroroutput: %v, %[2]T\n", atoi2, err2)
	}
}
