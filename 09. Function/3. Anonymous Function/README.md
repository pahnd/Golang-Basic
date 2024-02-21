Anonymous function (hàm ẩn danh) trong Go là một loại hàm không có tên, được khai báo trực tiếp trong mã chương trình mà không cần đặt tên cho nó. Hàm ẩn danh thường được sử dụng khi bạn cần viết một hàm đơn giản và không muốn định nghĩa nó trong một khai báo riêng biệt.

Để khai báo một hàm ẩn danh trong Go, bạn sử dụng cú pháp sau:


<code>
func(parameters) {
    // Thân hàm (function body)
}
</code>

Bạn có thể gán hàm ẩn danh cho một biến và gọi nó thông qua biến đó. Cũng có thể gọi hàm ẩn danh trực tiếp mà không cần gán nó vào biến.

Dưới đây là một ví dụ về hàm ẩn danh trong Go:


<code>

package main

import "fmt"

func main() {
    // Gọi hàm ẩn danh trực tiếp
    func() {
        fmt.Println("Đây là một hàm ẩn danh.")
    }()

    // Gán hàm ẩn danh cho biến và gọi nó thông qua biến
    add := func(a, b int) int {
        return a + b
    }

    result := add(3, 5)
    fmt.Println("Kết quả:", result)
}
</code>

Kết quả sẽ là:
<code>
Đây là một hàm ẩn danh.
Kết quả: 8
</code>


Trong ví dụ trên, chúng ta đã viết một hàm ẩn danh trực tiếp và gọi nó mà không cần gán cho một biến.

Tiếp theo, chúng ta đã khai báo một hàm ẩn danh với tên là add, nhận hai tham số a và b, và trả về tổng của chúng. Sau đó, chúng ta gán hàm ẩn danh này cho biến add. Bây giờ, biến add trở thành một function value, và chúng ta có thể gọi nó thông qua biến này như một hàm bình thường.

Hàm ẩn danh rất hữu ích trong các trường hợp mà bạn muốn viết một hàm đơn giản mà không cần định nghĩa nó trong một khai báo riêng biệt, hoặc khi bạn muốn tạo các function value linh hoạt và tái sử dụng.