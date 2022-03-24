package main

import "fmt"

func main() {
	var x interface{}
	var y interface{} = []int{3, 5}
	a := x == x
	b := x == y
	fmt.Println(a, b)
}
