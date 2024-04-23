package main

import "fmt"

/**
syntax:
	const identifier [type] = value

	Ex: const PI = 3.14159
**/

func main() {
	const HEIGHT int = 100
	const WIDTH int = 200
	area := HEIGHT * WIDTH
	fmt.Printf("value of area : %d", area)
}
