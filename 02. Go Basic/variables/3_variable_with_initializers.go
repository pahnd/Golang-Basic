package main

import "fmt"

//A var declaration can include initializers, one per variable.
var i, j int = 1, 2

func main() {
	//Nếu khai báo như trên type of variables sẽ được tự động define theo biến.
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
