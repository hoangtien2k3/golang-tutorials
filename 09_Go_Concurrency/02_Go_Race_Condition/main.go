package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// A race condition occurs in Go when two or more goroutines try to access the same resource.
// It may happen when a variable attempts to read and write the resource without any regard to other routines.

// Tình trạng dồn đuổi xảy ra trong Go khi hai hoặc nhiều goroutine cố gắng truy cập vào cùng một tài nguyên.
// Nó có thể xảy ra khi một biến cố gắng đọc và ghi tài nguyên mà không quan tâm đến các quy trình khác.

// Go Race Condition Example

var wait sync.WaitGroup
var count int

func increment(s string) {
	for i := 0; i < 10; i++ {
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)
	}
	wait.Done()
}

func main() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value ", count)
}
