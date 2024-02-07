package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

/*Biến khai báo without an explicit(rõ ràng) initial(ban đầu) value are given their zero value
The value is:
	0 for numeric types, false for boolean type and " " for string
*/
