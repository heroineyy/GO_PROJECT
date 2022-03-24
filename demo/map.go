package main

import "fmt"

func main(){
	m:= make(map[string] int)
	m["a"] = 1
	x,ok :=m["a"]
	fmt.Println(x,ok)

	delete(m,"a")
}
