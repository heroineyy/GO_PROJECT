package main

import "fmt"

//recover defer panic
func main() {
	fmt.Println("函数开始执行！")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("panic")
	return

}
