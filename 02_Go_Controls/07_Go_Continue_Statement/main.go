package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 1
	/* do loop execution */
	for a < 10 {
		if a == 5 {
			/* skip the iteration */
			a = a + 1
			continue
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}
