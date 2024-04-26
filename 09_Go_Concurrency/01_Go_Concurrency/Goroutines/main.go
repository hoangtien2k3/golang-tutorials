package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutines: A goroutine is a lightweight thread managed by the Go runtime.
// Ex: go f(x, y, z)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go fun1()
	go fun2()
	wg.Wait()
}

func fun1() {
	for i := 0; i < 10; i++ {
		fmt.Println("fun1,  ->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

func fun2() {
	for i := 0; i < 10; i++ {
		fmt.Println("fun2,  ->", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}
