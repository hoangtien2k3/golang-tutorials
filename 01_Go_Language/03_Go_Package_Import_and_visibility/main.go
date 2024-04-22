package main

import (
	"fmt"
	"os"
)

func main() {
	// Sử dụng hàm Getwd() từ thư viện os để lấy đường dẫn thư mục làm việc hiện tại
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Lỗi:", err)
		return
	}

	// In ra đường dẫn thư mục làm việc hiện tại
	fmt.Println("Thư mục làm việc hiện tại:", wd)
}
