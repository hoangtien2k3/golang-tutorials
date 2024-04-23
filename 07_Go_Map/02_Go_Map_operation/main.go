package main

/*
syntax:

	var map1 map[keytype]valuetype

Ex:

	var map1 map[int]string

*/

import "fmt"

func main() {

	// Go Map Insert and Update operation
	m := make(map[string]int)
	fmt.Println(m)
	m["Key1"] = 10
	m["Key2"] = 20
	m["Key3"] = 30
	fmt.Println(m)
	m["Key2"] = 555
	fmt.Println(m)

	/////////////

	// Go Map Delete operation
	m1 := make(map[string]int)
	m1["Key1"] = 10
	m1["Key2"] = 20
	m1["Key3"] = 30
	fmt.Println(m1)
	delete(m1, "Key3")
	fmt.Println(m1)

	//////////////

	// Go Map Retrieve Element
	// syntax: elem = m[key]
	m2 := make(map[string]int)
	m2["Key1"] = 10
	m2["Key2"] = 20
	m2["Key3"] = 30
	fmt.Println(m)
	v, ok := m2["Key2"] // truy cập tới 1 phần tử của map, thì Go sẽ trả về hai giá trị: giá trị của phần tử đó và một boolean cho biết phần tử đó có tồn tại trong map không.
	fmt.Println("The value:", v, "Present?", ok)
	i, j := m2["Key4"]
	fmt.Println("The value:", i, "Present?", j)
}
