package main

import "fmt"

type Walker interface {
	Walk(int) error
}

type Human struct {
	km int
}

func (h *Human) Walk(distance int) error {
	h.km += distance

	return nil
}

func main() {
	h := &Human{
		km: 0,
	}

	err := h.Walk(10)
	if err != nil {
		panic(err)
	}

	fmt.Println("Distance:", h.km)
	fmt.Println("Distance:", (*h).km)
	fmt.Printf("type: %T\naddress: %p\n%v\n", h, h, h)
}
