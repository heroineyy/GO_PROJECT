package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取传入的变量type,kind,值
	//1、先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp:", rTyp)

	//2.h获取到reflect.Value
	rVal := reflect.ValueOf(b)

	fmt.Println("rVal=", rVal)
}

func main() {

}
