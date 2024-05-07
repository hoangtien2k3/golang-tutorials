package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	n := 10000
	maxWorkers := 5

	wg := sync.WaitGroup{}

	queueCh := make(chan int, n)

	for i := 0; i < n; i++ {
		queueCh <- i
	}

	close(queueCh)

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func(count int) {
			for v := range queueCh {
				time.Sleep(time.Millisecond * 20)
				fmt.Printf("Worker %d is crawling we url %d\n\n", count, v)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

}
