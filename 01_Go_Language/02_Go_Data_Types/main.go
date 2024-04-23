package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%T %T %T %T\n", i, f, b, s)

	fmt.Printf("%v   %v      %v  %q     \n", i, f, b, s)

	///////
	var (
		name    string
		address string
		age     int
	)

	name = "Tien"
	address = "Thanh Hoa"
	age = 21
	fmt.Println(name, address, age)

	var name1, address1, age1 = "Tien 1", "Thanh Hoa 1", 21
	fmt.Println(name1, address1, age1)

	fmt.Println("==================")
	var myString string = "Tiến"

	runes := []rune(myString)
	fmt.Println(myString)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}

	fmt.Println("==================")
	var myRune rune = 'A'
	fmt.Printf("%T", myRune)

}
