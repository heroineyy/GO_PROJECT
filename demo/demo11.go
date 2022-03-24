package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

func main() {
	slice := make([]int, 5, 5) //len 5,cap 5 00000
	slice[0] = 1               //12000
	slice[1] = 2               //12000
	change(slice...)           //120003
	fmt.Println(slice)
	change(slice[0:2]...) //123
	fmt.Println(slice)
}
