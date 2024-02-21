package main

import (
	"fmt"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	x := -3.5
	f := MyFloat(-2.5)
	fmt.Println(f.Abs())
	b := MyFloat(x)
	fmt.Println(b)
	c := b.Abs()
	fmt.Println(c)
}
