package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func MapValues(values []int, mapFunc func(int) int) []int {
	var newValues []int
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func GenericMapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	result := MapValues([]int{1, 2, 3, 4}, func(n int) int {
		return n * 5
	})
	fmt.Printf("result: %+v\n", result)

	result1 := GenericMapValues([]int{1, 2, 3, 4}, func(n int) int {
		return n * 5
	})
	fmt.Printf("result: %+v\n", result1)

}
