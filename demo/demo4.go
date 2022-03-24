package main

//用chan做锁

import (
	"fmt"
	"time"
)

func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<-ch
}

func main() {
	pipeline := make(chan bool, 1)

	var x int
	for i := 0; i < 1000; i++ {
		go increment(pipeline, &x)
	}
	time.Sleep(time.Second)
	fmt.Println("x的值：", x)
}
