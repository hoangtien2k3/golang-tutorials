package main

import (
	"fmt"
)

// Struct
type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	x := person{firstName: "Tien", lastName: "Hoang", age: 21}

	fmt.Println(x)
	fmt.Println(x.firstName)
}
