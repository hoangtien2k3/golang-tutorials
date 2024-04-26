package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup // Khai báo WaitGroup

	// Thêm 3 Goroutines vào WaitGroup
	wait.Add(3)

	// Khởi chạy các Goroutines
	go func() {
		defer wait.Done() // Thông báo Goroutine đã hoàn thành khi kết thúc
		// Thực hiện công việc
		fmt.Println("Goroutine 1")
	}()

	go func() {
		defer wait.Done() // Thông báo Goroutine đã hoàn thành khi kết thúc
		// Thực hiện công việc
		fmt.Println("Goroutine 2")
	}()

	go func() {
		defer wait.Done() // Thông báo Goroutine đã hoàn thành khi kết thúc
		// Thực hiện công việc
		fmt.Println("Goroutine 3")
	}()

	// Chờ cho tất cả các Goroutines kết thúc
	wait.Wait()

	fmt.Println("All Goroutines are done")
}
