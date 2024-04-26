package main

/*
syntax base:

	type TypeName[T any] struct {
		// Trong đây bạn có thể sử dụng T như một kiểu dữ liệu chung
		Field T
	}

*/

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

}
