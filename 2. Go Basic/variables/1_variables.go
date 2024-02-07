package main

import "fmt"

//x int, y int or x, y int
//int sau cùng là kiểu dữ liệu trả về của hàm
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
