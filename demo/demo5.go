package main

// 使用 WaitGroup实现进程同步

import (
	"fmt"
	"sync"
)

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d:%d\n", x, i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()

}
