package main

// A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
// v, ok := <-ch
// `ok` is false if there are no more values to receive and the channel is closed.
// The loop for i := range c receives values from the channel repeatedly until it is closed.

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10) // buffer channel
	go fibonacci(cap(c), c) // cap: capacity của channel
	for i := range c {
		fmt.Println(i)
	}
}
