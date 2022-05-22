package main

import "fmt"

type Student struct {
	age  int
	Name string
}

func main() {
	ss := []Student{
		{
			Name: "xyy",
			age:  18,
		},
		{
			Name: "xss",
			age:  20,
		},
	}

	stu := []*Student{}
	for _, r := range ss {
		stu = append(stu, &r)
	}

	for _, tt := range stu {
		fmt.Println(tt.Name)
	}
}
