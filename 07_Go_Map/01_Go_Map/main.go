package main

/*
syntax:

	var map1 map[keytype]valuetype

Ex:

	var map1 map[int]string

*/

import "fmt"

func main() {
	x := map[string]int{"Kate": 28, "John": 37, "Raj": 20}
	fmt.Print(x)
	fmt.Println("\n", x["Raj"])
}
