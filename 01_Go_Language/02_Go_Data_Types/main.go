package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%T %T %T %T\n", i, f, b, s)

	fmt.Printf("%v   %v      %v  %q     \n", i, f, b, s)
}
