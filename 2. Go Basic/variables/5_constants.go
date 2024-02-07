package main

import "fmt"

const Pi = 3.14

func main() {
	cons World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rule?", Truth )
}
//Constants được khai báo giống như biến, nhưng vơi const keyword.
//Constants can be character, string, boolean, or numeric values.
//Constants cannnot be declared using the := syntax