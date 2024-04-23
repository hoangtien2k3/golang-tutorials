package main

import (
	"fmt"
	"reflect"
)

func main() {
	rune := 'A'
	fmt.Printf("%d \n", rune)
	fmt.Printf("%T \n", rune)
	fmt.Println(reflect.TypeOf(rune))
}
