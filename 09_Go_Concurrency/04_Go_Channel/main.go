package main

import (
	"fmt"
	"time"
)

// Channel là một tính năng mà qua đó cho phép các goroutines trao đổi dữ liệu với nhau.
// goroutine khi truyền dữ liệu vào channel thì phải có một goroutine khác lấy dữ liệu ra hoặc ngược lại
// nếu không các goroutine sẽ đi vào trạng thái “chờ”.
/*
syntax:
	ch := make(chan int)

	=> chúng được sử dụng với operator '<-'
	ch <- v    // Send v to channel ch.
	v := <-ch  // Receive from ch, and
           		// assign value to v.
*/

/*
Chỉ cho phép channel chỉ được phép truyền dữ liệu vào:

		send(c chan <- int)

Chỉ cho phép channel chỉ được phép đọc dữ liệu:

		receive(c <- chan int)

*/

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

// hàm send truyền giá trị của i vào channel c.
// Hàm receive nhận các giá trị từ channel c và assgin nó cho s rồi in ra màn hình.
func main() {
	c := make(chan int)
	go send(c)
	go receive(c)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("end")
}
