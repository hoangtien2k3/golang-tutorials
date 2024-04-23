package main

import "fmt"

func addItem(items int, list ...int) {
	// 100, 200, 300, 400, 500 -> []int {100, 200, 300, 400, 500}
	list = append(list, items)
	fmt.Println(list)
}

func change(list ...int) {
	list[0] = 999
}

func main() {
	addItem(1, 100, 200, 300, 400, 500)

	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	addItem(3, list...)

	change(list...)
	fmt.Println(list)

}
