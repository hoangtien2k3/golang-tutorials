package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Print(i, "\n")
	}

	for a := 0; a < 3; a++ {
		for b := 3; b > 0; b-- {
			fmt.Print(a, " ", b, "\n")
		}
	}

}
