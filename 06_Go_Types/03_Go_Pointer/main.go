package main

/*
syntax:
	var var_name *var-type
*/

/*
1. Khai báo con trỏ cho một kiểu dữ liệu cụ thể:
	var ptr *int  // Con trỏ đến kiểu dữ liệu int


2. Khai báo và gán con trỏ với giá trị của một biến:
	var x int = 10
	var ptr *int = &x  // ptr trỏ đến biến x


3. Khai báo con trỏ không khởi tạo (nil pointer):
	var ptr *int  // Khai báo con trỏ không trỏ đến bất kỳ biến nào (nil)


4. Sử dụng từ khóa new để cấp phát bộ nhớ cho một biến và trả về một con trỏ đến biến đó:
	ptr := new(int)  // ptr là con trỏ đến một biến kiểu int được cấp phát bộ nhớ

*/

import (
	"fmt"
)

func changeX(x *int) {
	*x = 0
}

func changePtr(ptr *int) {
	*ptr = 10
}

func main() {
	x := 10
	changeX(&x)
	fmt.Println("x = ", x)

	////////

	ptr := new(int)
	fmt.Println("Before change ptr", *ptr)
	changePtr(ptr)
	fmt.Println("After change ptr", *ptr)

}
