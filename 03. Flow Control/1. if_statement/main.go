package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// converting string to int:
	i, err := strconv.Atoi("45")

	// error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)

	}

	// simple (short) statement ->  the same effect as the above code
	// i and err are variables scoped to the if statement only
	// The if statement can start with a short statement to execute before the condition.
	if i, err := strconv.Atoi("34"); err == nil {
		fmt.Println("No error. i is ", i)
	} else {
		fmt.Println(err)
	}
	ValidateUser()
	Age()
	VowelCons_1()
	VowelCons_2()
}

func ValidateUser() {
	const (
		usage    = "Usage: [username] [password]"
		errUser  = "Access denied for %q. \n"
		errPwd   = "Invaild password for %q. \n"
		accessOK = "Access granted to %q. \n"

		user, user2 = "jack", "inanc"
		pass, pass2 = "1888", "1879"
	)
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := os.Args[1], os.Args[2]

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p == pass {
		fmt.Printf(accessOK, u)
	} else if u == user2 && p == pass {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}

	return
}

func Age() {
	age := 10
	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}
	return
}

func VowelCons_1() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}
	s := args[1]
	if s == "a" || s == "e" || s == "o" || s == "u" {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}
	return
}

func VowelCons_2() {
	args := os.Args

	if len(args) != 2 || len(agrs[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	s := args[1]
	if strings.IndexAny(s, "aeiot") != -1 {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consoant.\n", s)
	}
	return
}
