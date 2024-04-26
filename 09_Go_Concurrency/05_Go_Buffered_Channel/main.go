package main

import (
	"fmt"
	"time"
)

// goroutine khi truyền dữ liệu vào channel thì phải có một goroutine khác lấy dữ liệu ra hoặc ngược lại
// nếu không các goroutine sẽ đi vào trạng thái “chờ”.
// => goroutine không chờ nữa bằng cách dùng các buffered channel
// Buffered Channel tức là channel đó giới hạn dữ liệu vào.

func send(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("send %v to channel\n", i)
		c <- i // send i to channel
	}
}

func receive(c chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		s := <-c // receive channel
		fmt.Println(s)
	}
}

func main() {
	// Hàm send liên tục đưa dữ liệu vào channel c mà không đợi hàm receive lấy nó ra.
	c := make(chan int, 5) // ta thay c := make(chan int) thành c := make(chan int, 5)
	go send(c)
	go receive(c)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("end")
}
