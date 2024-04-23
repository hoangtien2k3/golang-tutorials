package main

import (
	"fmt"
)

func main() {
	odd := [6]int{2, 4, 6, 8, 10, 12}
	var s []int = odd[1:4]
	fmt.Println(s)

	///////////

	names := [4]string{
		"John",
		"Jim",
		"Jack",
		"jen",
	}
	fmt.Println(names)
	slice1 := names[0:2]
	slice2 := names[1:3]
	fmt.Println(slice1, slice2)
	slice2[0] = "ZZZ"
	fmt.Println(slice1, slice2)
	fmt.Println(names)

	//////////
	stru := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, true},
		{5, false},
		{6, true},
	}
	fmt.Println(stru)
}
