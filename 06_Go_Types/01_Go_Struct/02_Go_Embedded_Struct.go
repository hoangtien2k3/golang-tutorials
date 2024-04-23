package main

import (
	"fmt"
)

type personEx struct {
	fname string
	lname string
}

type employee struct {
	personEx
	empId int
}

func (p personEx) details() {
	fmt.Println(p, " "+" I am a person")
}

func (e employee) details() {
	fmt.Println(e, " "+"I am a employee")
}

func main() {
	p1 := personEx{fname: "hoangtien", lname: "2k3"}
	p1.details()
	e1 := employee{personEx: personEx{"John", "Ponting"}, empId: 11}
	e1.details()
}
