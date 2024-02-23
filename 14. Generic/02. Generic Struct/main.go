package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	// Data interface{} không nên sử dụng
	Data T
}

func main() {
	u := User[int]{
		ID:   0,
		Name: "",
		Data: 3,
	}
	fmt.Printf("user: %+v/n", u)
	u1 := User[string]{
		ID:   0,
		Name: "",
		Data: "abc",
	}
	fmt.Printf("user: %+v/n", u1)
}
