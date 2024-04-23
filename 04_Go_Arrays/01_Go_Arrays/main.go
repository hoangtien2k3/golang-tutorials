package main

/*
syntax:
	var identifier [len]type
*/

import "fmt"

func main() {
	var x [5]int
	var i, j int
	for i = 0; i < 5; i++ {
		x[i] = i + 10
	}
	for j = 0; j < 5; j++ {
		fmt.Printf("Element[%d] = %d\n", j, x[j])
	}

	fmt.Println("================\n")

	// value type
	array3 := [...]string{"VN", "FRANCE", "CHINA"}
	fmt.Println(array3)

	copyArray3 := array3 // copyArray3 không phải là reference type mà là value mới
	copyArray3[0] = "CANADA"

	fmt.Println(array3)
	fmt.Println(copyArray3)

	// C1: duyệt vòng lặp theo index:
	for i := 0; i < len(array3); i++ {
		fmt.Println(array3[i])
	}

	// C2: dùng for range
	for index, value := range array3 {
		fmt.Printf("i = %d - value = %s", index, value)
		fmt.Println()
	}

}
