package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Comprison Operators
	s1 := "Fish"
	s2 := "fish"
	s3 := "5" // String

	i1 := 123
	i2 := 321

	fmt.Println(s1 == s2)                  // False
	fmt.Println(strings.ToLower(s1) == s2) // True

	// fmt.Println(s1 == i1) // Error

	fmt.Println(i1 == i2) // False
	fmt.Println(i1 != i2) // True
	fmt.Println(i1 >= i2) // False
	fmt.Println(i1 > i2)  // False
	fmt.Println(i1 <= i2) // True
	fmt.Println(i1 < i2)  // True

	// Arithmetic Operators
	fmt.Println(s1 + s2)

	// fmt.Println(s3 + i1) // Error
	fmt.Println(s3 + strconv.Itoa(i1))

	// fmt.Println(i1 + s3) // Error
	atoi, err := strconv.Atoi(s3)
	if err == nil {
		fmt.Println(i1 + atoi)
	}

	atoi2, err2 := strconv.Atoi(s3)
	_ = err2 // Hack not to be used in real world to remove "assigned and never used for error"
	fmt.Println(i1 + atoi2)

	i3 := 10
	i4 := 5

	fmt.Println(i3 + i4)
	fmt.Println(i3 - i4)
	fmt.Println(i3 * i4)
	fmt.Println(i3 / i4)
	fmt.Println(i3 % i4)

	i3++
	fmt.Println(i3)

	i3--
	fmt.Println(i3)

}
