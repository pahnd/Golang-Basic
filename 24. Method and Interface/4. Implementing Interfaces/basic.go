package main

import (
	"fmt"
)

type Speak interface {
	S()
}

type Human struct {
	name string
}

func (human *Human) S() {
	fmt.Println(human.name)
}

type Radio string

func (r Radio) S() {
	fmt.Println(r)
}

func main() {
	var i Speak

	i = &Human{"Jack"}
	describe(i)
	i.S()

	i = Radio("Radio is speaking")
	describe(i)
	i.S()
}

func describe(s Speak) {
	fmt.Printf("(Speaking: %v, %T)\n", s, s)
}
