package main

import "fmt"

type myType struct {
	i int
}

func pointerFuncInt(i *int) {
	*i += 10
}

func pointerFuncStruct1(m *myType) {
	m.i += 10
}

func (m *myType) pointerFuncStruct2() {
	// Doesn't need (*m).i
	m.i += 10
	(*m).i += 10
}

func pointerFuncSliceString(s []string) {
	for obj := range s {
		s[obj] += " has changed"
	}
}

func main() {
	a := 10
	pointerFuncInt(&a)
	fmt.Println(a)

	b := myType{i: 10}
	pointerFuncStruct1(&b)
	fmt.Println(b)

	// Doesn't need (&b).
	b.pointerFuncStruct2()
	(&b).pointerFuncStruct2()
	fmt.Println(b)

	s := []string{"1", "2"}
	pointerFuncSliceString(s)
	fmt.Println(s)
}
