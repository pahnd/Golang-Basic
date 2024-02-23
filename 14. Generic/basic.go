package main

import "fmt"

type Num interface {
	int | int8 | int16 | float32 | float64
}

type UserID int

func AddInt(a, b int) int {
	return a + b
}

func AddFloat(a, b float64) float64 {
	return a + b
}

func Add[T Num](a T, b T) T {
	return a + b
}

func Add1[T int | float64](a T, b T) T {
	return a + b
}

func Add2[T ~int | float64](a T, b T) T { //~int allow all type int
	return a + b
}

func main() {
	result1 := AddInt(3, 5)
	fmt.Printf("result: %+v\n", result1)

	//	Error because AddInt receive only Inter value
	//	result_1 := AddInt(1.1, 2.2)

	result2 := AddFloat(1.1, 2.2)
	fmt.Printf("result: %+v\n", result2)

	result3 := Add(3, 5)
	fmt.Printf("result: %+v\n", result3)

	result4 := Add(1.1, 2.2)
	fmt.Printf("result: %+v\n", result4)

	a := UserID(1)
	b := UserID(2)
	// Add1(a,b) Fail UserID does not implement int|float64
	Add2(a, b) // Ok

}
