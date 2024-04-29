package main

import (
	"fmt"
	"time"
)

// Channels are a typed conduit through which you can send and receive values with the channel operator: '<-'
/*
syntax:
	ch <- v    	// Send v to channel ch.
	v := <-ch  	// Receive from ch, and
           		// assign value to v.

syntax: create channels
	ch := make(chan int)
*/
// Like maps and slices, channels must be created before use:
// ch := make(chan int)
// By default, sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.
// The example code sums the numbers in a slice, distributing the work between two goroutines.
// Once both goroutines have completed their computation, it calculates the final result.

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

/////////
//package main
//
//import "fmt"
//
//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	c <- sum // send sum to c
//}
//
//func main() {
//	s := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go sum(s[:len(s)/2], c)
//	go sum(s[len(s)/2:], c)
//	x, y := <-c, <-c // receive from c
//
//	fmt.Println(x, y, x+y)
//}
