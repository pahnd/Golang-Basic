package main

import (
	"fmt"
)

func multi(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := multi("Hello", "World")
	fmt.Println(a, b)
}

// A function can return any number of results.
// The swap function returns two strings.
