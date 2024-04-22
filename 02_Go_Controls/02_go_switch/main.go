package main

import "fmt"

func main() {
	fmt.Print("Enter Number: ")
	var input int
	fmt.Scanln(&input)

	switch input {
	case 10:
		fmt.Print("the value is 10")
	case 20:
		fmt.Print("the value is 20")
	case 30:
		fmt.Print("the value is 30")
	case 40:
		fmt.Print("the value is 40")
	default:
		fmt.Print(" It is not 10,20,30,40 ")
	}
}
