package main

import "fmt"

func main() {
	// Gọi một hàm để xử lý dữ liệu
	err := process()

	// Kiểm tra lỗi trả về từ hàm process
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Hàm process trả về một lỗi nếu có lỗi xảy ra
func process() error {
	// Giả sử xảy ra một tình huống không thể xử lý
	// Ví dụ: chia cho 0
	// Trong trường hợp này, chúng ta sẽ gây ra một panic
	panic("Something went wrong!")
}
