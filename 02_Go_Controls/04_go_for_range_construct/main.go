package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

/**
syntax:

	for ix, val := range coll {

	}

Trong đó:
	+ index là chỉ số của phần tử trong collection (hoặc là key nếu collection là map).
	+ value là giá trị của phần tử tại chỉ số (hoặc là value tương ứng nếu collection là map).
	+ collection là cấu trúc dữ liệu mà bạn muốn lặp qua các phần tử.

**/

func main() {

	// 1. Lặp qua slice hoặc array:
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 2. Lặp qua map:
	ages := map[string]int{"John": 30, "Alice": 25, "Bob": 35}
	for key, value := range ages {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// 3. Lặp qua chuỗi (string):
	message := "Hello, world!"
	for index, char := range message {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// 4. Lặp qua channel:
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for num := range ch {
		fmt.Println("Received:", num)
	}

	// 5. Lặp qua file hoặc đường ống (io.Reader):
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

}
