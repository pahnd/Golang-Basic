package main

import "fmt"

// define dữ liệu return sẽ trả về x và y
//A return statement without arguments returns the named return values. This is known as a "naked" return.
//Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
func split(sum int) (x, y int) {
	fmt.Println("x có giá trị là: ", sum*10)
	fmt.Println("y có giá trị là: ", sum-10)
	x = sum * 10
	y = sum - 10
	return
}
func main() {
	fmt.Println(split(10))
}
