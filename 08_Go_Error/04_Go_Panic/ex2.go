package main

import "fmt"

func main() {
	fmt.Println("Calling x from main.")
	x()
	fmt.Println("Returned from x.")
}

func x() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in x", r)
		}
	}()
	fmt.Println("Executing x...")
	fmt.Println("Calling y.")
	y(0)
	fmt.Println("Returned normally from y.")
}

func y(i int) {
	fmt.Println("Executing y....")
	if i > 2 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in y", i)
	fmt.Println("Printing in y", i)
	y(i + 1)
}
