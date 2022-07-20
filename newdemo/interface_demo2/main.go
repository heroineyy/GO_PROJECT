//接口使用的注意事项
package main

import "fmt"

//结构体变量实现接口
type a interface {
	test()
}

type A struct {
	name string
}

func (a1 A) test() {
	fmt.Println("hello yy")
}

//整型变量实现接口
type B int

func (b1 B) test() {
	fmt.Println("hello ", b1)
}

func main() {
	var a1 A
	var aInterface a = a1
	aInterface.test()

	var b1 B = 10
	var bInterface a = b1
	bInterface.test()
}
