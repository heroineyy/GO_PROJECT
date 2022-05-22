package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//声明一个student结构体
type Student struct {
	name  string
	age   int
	score int
}

//声明一个student结构体切片类型
type studentSlice []Student

//实现Interface接口
func (st studentSlice) Len() int {
	return len(st)
}

//less这个方法就是决定你使用什么方法排序
func (st studentSlice) Less(i, j int) bool {
	return st[i].age > st[j].age
}

func (st studentSlice) Swap(i, j int) {
	st[i], st[j] = st[j], st[i]
}

func main() {
	//声明一个结构体切片
	var students studentSlice
	for i := 0; i < 10; i++ {
		student := Student{
			name:  fmt.Sprintf("student%d", rand.Intn(100)),
			age:   rand.Intn(100),
			score: rand.Intn(100),
		}
		//将student append到student切片
		students = append(students, student)
	}
	//看看排序前的顺序
	for _, v := range students {
		fmt.Println(v)
	}
	fmt.Println("排序后-----")

	//调用sort.sort
	sort.Sort(students)
	for _, v := range students {
		fmt.Println(v)
	}

}
