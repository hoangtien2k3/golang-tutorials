package main

import "fmt"

func main() {
	var a int = 1
	for a < 10 {
		fmt.Print("Value of a is ", a, "\n")
		a++
		if a > 5 {
			/*
				terminate the loop using break statement
			*/
			break
		}
	}
}
