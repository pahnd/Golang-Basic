package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	Convert()
	FleetToMeter()
}

func Convert() {
	age := os.Args[1]

	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Printf("SUCCESS: Converted %q to %d. \n", age, n)
	return
}

func FleetToMeter() {
	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("ERROR: %q is not the number ", arg)
		return
	}
	meters := feet * 0.3048
	fmt.Printf("SUCCESS: %g feet is %g meters.\n", feet, meters)
	return
}
