package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	var f float64 = 6.44
	var str1 string = "101"
	var str2 string = "10.123"

	fmt.Println(float64(i))
	fmt.Println(int(f))

	newInt, _ := strconv.ParseInt(str1, 0, 64)
	fmt.Println(newInt)

	newfloat, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		fmt.Println("Lỗi:", err)
	} else {
		fmt.Println("newfloat:", newfloat)
	}

}
