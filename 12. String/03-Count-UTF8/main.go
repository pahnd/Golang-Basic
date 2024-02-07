package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "Nguyệt"
	length := len(text)
	lengthUTF8 := utf8.RuneCountInString(text)
	fmt.Println(length)
	fmt.Println(lengthUTF8)
}
