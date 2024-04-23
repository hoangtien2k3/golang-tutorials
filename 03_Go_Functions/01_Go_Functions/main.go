package main

/**
Go Functions:

	+ Normal functions with an identifier
	+ Anonymous or lambda functions
	+ Method (A function with a receiver)
**/

import "fmt"

type Employee struct {
	fname string
	lname string
}

func (emp Employee) fullname() {
	fmt.Println(emp.fname + " " + emp.lname)
}
func main() {
	e1 := Employee{"Hoang", "Tien"}
	e1.fullname()
}
