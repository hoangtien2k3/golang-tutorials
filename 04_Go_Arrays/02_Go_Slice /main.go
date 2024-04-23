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

	//////////

	var array = [4]int{1, 2, 3, 4}
	slice3 := array[:]
	fmt.Println(slice3)

	slice4 := array[2:]
	fmt.Println(slice4)

	slice5 := array[:3]
	fmt.Println(slice5)

	///////////////
	// Tạo một slice từ một slice khác

	var slice6 = []int{1, 2, 3}
	slice7 := slice6
	fmt.Println(slice7)

	// slice -> reference type
	var array12 = [5]int{1, 2, 3, 4, 5}
	slice8 := array12[:]
	fmt.Println(slice8)

	slice8[0] = 777
	fmt.Println(array12)

	// length và capacity của slice
	countries := [...]string{"VN", "US", "CANADA", "FRANCE", "CHINA"}
	slice10 := countries[2:3]
	fmt.Println(slice10)

	fmt.Printf("Length slice10: %d", len(slice10))     // số lượng phần tử của slice
	fmt.Printf("\nCapacity slice10: %d", cap(slice10)) // số lươợng phần tử của underflying array: vị trí bắt đầu của slice trong mảng trở đi

	///////////////////

	// make, copy, append

	// 1. make: khởi taạo một slice bằng make
	slice11 := make([]int, 2)
	fmt.Println("\nSlice11: ", slice11)
	fmt.Println("Length slice11: ", len(slice11))
	fmt.Printf("\nCapacity slice11: %d", cap(slice11))

	// 2. append: thêm phần tử vào slice
	slice12 := make([]int, 2)
	slice12 = append(slice12, 100)
	fmt.Println("\nSlice12: ", slice12)

	// 3. copy
	src := []string{"A", "B", "C", "D", "E", "F"}
	dest := make([]string, 2)
	number := copy(dest, src)
	fmt.Printf("Copied %d elements from dst to src %s\n", number, dest)
	fmt.Println(dest)

	// delete item with index 1
	src = append(src[:1], src[2:]...) // nối 2 slice lại  với nhau
	fmt.Println(src)
}
