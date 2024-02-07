package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args:", os.Args) // os.Args is slice of strings ([]string)

	// accessing command line arguments using indexes
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st Argument:", os.Args[1])
	fmt.Println("2nd Argument:", os.Args[2])
	fmt.Println("No. of items inside os.Args:", len(os.Args))
}

// To test this code run: go run main.go linux window mac 100
// Tất cả tham số được truyền vào từ arg đều dạng string. Nếu muốn sử dụng interger thì phải converse strconv.ParseFload(os.Args[1], 64 )
