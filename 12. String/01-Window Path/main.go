package main

import "fmt"

func main() {
	//this one uses a raw string literal
	path := `c:\program files\New Folder\_text.txt
	c:\program files\New Folder\_text_1.txt
	`
	fmt.Print(path)
}
