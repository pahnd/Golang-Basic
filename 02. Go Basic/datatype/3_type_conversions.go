package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt((float64(x*x + y*y))) //math.Sqrt tính căn bậc 2. float64 value in argument to math.Sqrt
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
