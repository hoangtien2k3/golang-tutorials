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

// single return value
func Chao(name string) string {
	result := fmt.Sprintf("Hello: %s", name)
	return result
}

// multiple return value
func Hight(w int, h int) (int, int) {
	return w, h
}

// multiple return value
func ReacHight(w int, h int) (int, int, int) {
	area := w * h
	return w, h, area
}

// named return value
func ReacInfo(w int, h int) (weight int, hight int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}

func main() {
	e1 := Employee{"Hoang", "Tien"}
	e1.fullname()

	fmt.Println("==================\n")

	result := Chao("hoangtien2k3")
	fmt.Println(result)

	fmt.Println("==================\n")

	w, h := Hight(50, 170)
	fmt.Println(w, h)

	fmt.Println("==================\n")
	w1, h2, area := ReacHight(50, 170)
	fmt.Println(w1, h2, area)

	fmt.Println("==================\n")
	w2, h2, isSquare := ReacInfo(50, 170)
	if isSquare {
		fmt.Println("This is square")
	} else {
		fmt.Println("This is not square")
		fmt.Println(w2)
		fmt.Println(h2)
	}
}
