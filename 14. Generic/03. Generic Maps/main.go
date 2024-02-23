package main

//type CustomMap map[string]int
//type CustomMap map[interface{}]int
type CustomMap[T comparable, V int | string] map[T]V

//comparable là bất cứ giá trị nào có thể so sánh. Chú ý: Kiểu dữ liệu nào được sử dụng để sử dụng cho giá trị key của map => comparable
//https://levelup.gitconnected.com/comparable-in-golang-d7b080e49443
func main() {
	m := make(CustomMap[int, string])
	m[3] = "3"
}
