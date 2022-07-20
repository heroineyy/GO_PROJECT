package main

type T struct{}

func (*T) foo() {
}

func (T) bar() {
}

func (T) bar1() {
}

type S1 struct {
	*T
}

func main() {
	s := S1{}
	_ = s.foo //Q1
	s.foo()
	_ = s.bar //Q2
}
