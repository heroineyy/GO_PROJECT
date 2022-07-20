package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

type college interface {
	English()
}

func (s student) study() {
	fmt.Sprintf("student %s love study", s.name)
}

type collegestudent struct {
	student
}

func (cS collegestudent) English() {
	fmt.Printf("college student %s lean english\n", cS.name)
}

func main() {
	var cSt collegestudent
	var colStu college = cSt
	colStu.English()

}
