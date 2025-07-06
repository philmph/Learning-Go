package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func newCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func main() {
	m1 := newCustomMap[int, string]()
	m1.Insert(1, "Hello")
	m1.Insert(2, "World")

	for k, v := range m1.data {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}
}
