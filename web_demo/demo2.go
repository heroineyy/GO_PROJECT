package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析url传递的参数
	fmt.Println(r.Form) //将信息输出到服务器
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "helo xyy")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

func main() {
	//http.HandleFunc("/", sayhelloName1)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("listenandserver:", err)
	}
}
