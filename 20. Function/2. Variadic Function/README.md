Hàm biến thể (variadic function) trong Go là một loại hàm cho phép chấp nhận một số lượng biến đổi từ đối số. Dấu ba chấm (...) được sử dụng để chỉ định một tham số biến thể trong danh sách tham số của hàm. Điều này cho phép bạn truyền vào hàm một số lượng đối số có thể là 0 hoặc nhiều đối số cùng một kiểu dữ liệu.

Dưới đây là một ví dụ về hàm biến thể trong Go:

go

<code>
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Tổng:", result)

	// Truyền không có đối số vào hàm sum
	result = sum()
	fmt.Println("Tổng:", result)

	// Truyền nhiều đối số vào hàm sum
	result = sum(10, 20, 30)
	fmt.Println("Tổng:", result)
}
</code>

Kết quả:

makefile
Copy code
Tổng: 15
Tổng: 0
Tổng: 60
Trong ví dụ trên, chúng ta đã định nghĩa một hàm biến thể có tên sum, cho phép nhận một số lượng không xác định các đối số kiểu int. Bên trong hàm, chúng ta sử dụng một vòng lặp for để duyệt qua các đối số được truyền vào và tính tổng chúng.

Trong hàm main, chúng ta thể hiện cách sử dụng hàm sum với các số lượng đối số khác nhau. Chúng ta gọi sum với 5, 0, và 3 đối số tương ứng. Dù số lượng đối số khác nhau, hàm sum có thể xử lý tất cả các trường hợp này nhờ tham số biến thể.

Hàm biến thể rất hữu ích khi bạn muốn tạo một hàm có thể được gọi với số lượng đối số linh hoạt mà không cần biết trước số lượng đối số sẽ được truyền vào.