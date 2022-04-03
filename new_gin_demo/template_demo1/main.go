package main

import (
	"fmt"
	"html/template"
	"net/http"
)

///首字母一定要大写
type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:%v", err)
		return
	}
	//渲染模板
	u1 := User{
		Name:   "xyy",
		Gender: "女",
		Age:    18,
	}

	m1 := map[string]interface{}{
		//map不用首字母大写
		"Name":   "小王子",
		"Gender": "男",
		"Age":    18,
	}
	hobbyList := []string{
		"run",
		"read",
		"coding",
	}

	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server start failed, err:%v", err)
		return
	}
}
