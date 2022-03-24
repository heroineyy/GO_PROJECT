package main

//关闭的空chan向其接受数据不报错，但是有值的chan就会报错
//我明白了，一个无缓存的chan有值就会想尽一切办法出去
import "fmt"

func main() {
	ch := make(chan int)
	ch <- 3
	close(ch)
	d := <-ch
	fmt.Println(d)
}
