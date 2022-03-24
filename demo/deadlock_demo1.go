package main

import "fmt"

func main() {
	pipelines := make(chan string)
	go func() {
		pipelines <- "hello world"
		pipelines <- "hello China"
		close(pipelines)
	}()
	for data := range pipelines {
		fmt.Println(data)
	}
}
