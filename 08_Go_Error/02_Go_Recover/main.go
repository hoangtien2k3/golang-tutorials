package main

import (
	"fmt"
)

// Defer, Panic, and Recover
func main() {
	fmt.Println(SaveDivide(10, 0))
	fmt.Println(SaveDivide(10, 10))
	fmt.Println(SaveDivide(20, -20))
}

func SaveDivide(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	quotient := num1 / num2
	return quotient
}
