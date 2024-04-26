package main

import (
	"fmt"
)

// defer là một keyword được sử dụng để lên lịch trình cho việc thực thi một hàm sau khi một hàm khác đã hoàn thành.
// Deferred function: Hàm được lên lịch trình bởi defer được gọi là "deferred function".
// LIFO (Last In, First Out): Các deferred function được thực thi theo thứ tự ngược lại so với thứ tự mà chúng được gọi.
func main() {
	defer print1("Hi...") // Lên lịch trình in ra "Hi..." sau khi hàm main kết thúc
	print2("there")       // In ra "there" trước khi deferred function được thực hiện

	// hiểu đơn giản: hàm khác thực thi xong thì hàm định nghĩa 'defer' mới được gọi
	// hàm main thực thi xong và in ra: there
	// hàm print1 mới thực thi: Hi...
}

func print1(s string) {
	fmt.Println(s)
}

func print2(s string) {
	fmt.Println(s)
}
