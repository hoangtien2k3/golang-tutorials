package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 10
	/* check the boolean condition */
	if a%2 == 0 {
		/* if condition is true then print the following */
		fmt.Printf("a is even\n")
	} else {
		/* if condition is false then print the following */
		fmt.Printf("a is odd\n")
	}

	fmt.Printf("value of a is : %d\n", a)
}
