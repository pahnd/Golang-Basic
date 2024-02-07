/////////////////////////////////
// Scopes in Go

/////////////////////////////////

// There are 3 Scopes:
//	- File Scope
// 	- Package Scope
// 	- Block (local)  Scope

package main

// import statements are file scoped
import (
	"fmt"
	"os"
	"strconv"

	// import "fmt" -> error, within the same scope, unique names

	// importing as another name (alias) is permitted
	f "fmt"
)

// variables or constant declared outside any function are package scoped
const done = false

func main() { // package scoped

	// block scoped: visible until the end of the block "}"
	var task = "Running:"
	fmt.Println(task, done) // => Running: false (this is done from package scope)
	f.Println("Bye bye!")

	// names must be unique only within the same scope
	const done = true                        // local scoped
	fmt.Printf("done in main(): %v\n", done) // => done in main(): true
	f1()
	Statement_If()
}

func f1() {
	fmt.Printf("done in f(): %v\n", done) //this is done from package scope
}

func Statement_If() {
	if a := os.Args; len(a) != 2 {
		// can access only: a variable
		fmt.Println("Give me a number.")
	}
	else if n, err := strconv.Atoi(a[1]); err != nil {
		// can access: a, n, err variables
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		// all the variables in the if statement
		fmt.Printf("%s * 2 %d\n", a[1], n*2)
	}
	return
}