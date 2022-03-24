package main

import "fmt"

func f1(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f1 func()

	defer f1()
	f1 = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f1(3))
}
