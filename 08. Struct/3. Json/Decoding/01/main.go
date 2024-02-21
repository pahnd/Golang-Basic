package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms"`
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user
	// Why do you need to pass a pointer to the Unmarshal function?

	//: Otherwise, it would not be able to update the given value.
	// It's because, every value in Go is passed by value.
	// So a function can only change the copy, not the original value.
	// However, through a pointer, a function can change the original value.
	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["write"]:
			fmt.Print(" can write.")
		}
		fmt.Println()
	}
}
