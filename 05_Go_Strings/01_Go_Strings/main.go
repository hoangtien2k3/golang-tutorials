package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

	fmt.Println("==============")

	str := "I love my country"
	fmt.Println(len(str))

	fmt.Println("==============")

	str1 := "india"

	// ToLower() , ToUpper()
	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.ToLower(str1))

	// HasPrefix()
	fmt.Println(strings.HasPrefix(str1, "IN"))

	// HasSuffix()
	fmt.Println(strings.HasSuffix(str1, "IA"))

	// Join()
	var arr = []string{"a", "b", "c", "d"}
	fmt.Println(strings.Join(arr, "*"))

	// Repeat()
	var str12 = "New "
	fmt.Println(strings.Repeat(str12, 4))

	// Contains()
	str13 := "Hi...there"
	fmt.Println(strings.Contains(str13, "th"))

	// Index()
	str14 := "Hi...there"
	fmt.Println(strings.Index(str14, "th"))

	// Count()
	str15 := "Hi...there"
	fmt.Println(strings.Count(str15, "e"))

	// Replace()
	str16 := "Hi...there"
	fmt.Println(strings.Replace(str16, "e", "Z", 2))

	// Split()
	str17 := "I,love,my,country"
	var arr17 []string = strings.Split(str17, ",")
	fmt.Println(len(arr17))
	for i, v := range arr17 {
		fmt.Println("Index : ", i, "value : ", v)
	}

	// Compare()
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))

	// Trim()
	fmt.Println(strings.TrimSpace(" \t\n I love my country  \n\t\r\n"))

	// ContainsAny()
	fmt.Println(strings.ContainsAny("Hello", "A"))
	fmt.Println(strings.ContainsAny("Hello", "o & e"))
	fmt.Println(strings.ContainsAny("Hello", ""))
	fmt.Println(strings.ContainsAny("", ""))

}
