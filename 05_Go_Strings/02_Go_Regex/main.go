package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(".com")
	fmt.Println(re.FindString("google.com"))
	fmt.Println(re.FindString("abc.org"))
	fmt.Println(re.FindString("fb.com"))

	//////

	rex := regexp.MustCompile(".com")
	fmt.Println(rex.FindStringIndex("google.com"))
	fmt.Println(rex.FindStringIndex("abc.org"))
	fmt.Println(rex.FindStringIndex("fb.com"))

	//////

	re1 := regexp.MustCompile("f([a-z]+)ing")
	fmt.Println(re1.FindStringSubmatch("flying"))
	fmt.Println(re1.FindStringSubmatch("abcfloatingxyz"))
}
