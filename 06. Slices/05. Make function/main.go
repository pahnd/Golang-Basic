// Use the "make" when you want to optimize your code or pass a limited size slice to some function (for example: checkout io.Read )
// Sử dụng "make" khi bạn muốn tối ưu hóa mã của mình hoặc chuyển một slice có kích thước giới hạn cho một số chức năng (ví dụ: kiểm tra io.Read )
// "make" can preallocate a backing array with the given length and/ỏ capacity
// "make" có thể phân bổ trước một mảng sao lưu với độ dài và/hoặc dung lượng nhất định

package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10
	withoutMake()
	makeLength()
	makeLengthCapacity()
}

func withoutMake() {
	tasks := []string{"jump", "run", "read"}

	var upTasks []string
	s.Show("upTasks", upTasks)
	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}
	return
}

func makeLength() {
	tasks := []string{"jump", "run", "read"}

	upTasks := make([]string, len(tasks))
	s.Show("upTasks", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}
}

func makeLengthCapacity() {
	tasks := []string{"jump", "run", "read"}

	upTasks := make([]string, 0, len(tasks))
	s.Show("upTasks", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}
}
