package main

import "fmt"

func main() {
	//可以
	//	te1 = &Person{
	//		Name: "小马",
	//		Age:  10,
	//	}
	//	te1.setName()调用 *T  T类型的方法
	var te1 myPerson //接口类型
	//不能直接定义接口，通过接口调用方法，必须实例化，指定接口的类型，即用什么实现的接口
	te1 = &Person{
		Name: "小马",
		Age:  10,
	}
	te1.setPerAge()
	fmt.Println(te1)

	//只能调用 T 类型的方法
	var people myPeople
	people = Person{
		Name: "小河",
		Age:  5,
	}
	people.setAge()
	fmt.Println(people)
}

// Person 定义结构体
type Person struct {
	Name string
	Age  int
}

//定义接口
type myPerson interface {
	setName()
	setPerAge()
}

//定义接口
type myPeople interface {
	setAge()
	//setPeopleName()
}

//person实现接口myperson
func (p *Person) setName() {
	p.Name = "小呼"
}
func (p Person) setPerAge() {
	fmt.Println("输出年龄:", p.Age)
}

//person实现接口myPeople
func (p Person) setAge() {
	p.Age = 18
}

//func (p *Person) setPeopleName() {
//    p.Name = "测试"
//}
