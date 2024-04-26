package main

import "os"

func main() {
	panic("Error Situation")
	_, err := os.Open("/tmp/file")
	if err != nil {
		panic(err)
	}
}
