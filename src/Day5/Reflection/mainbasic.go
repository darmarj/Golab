package main

import (
	"fmt"
	"reflect"
)

func main() {
	//反射操作: 通过反射，可以获得一个接口类型变量 类型和数据
	var x float64 = 3.4

	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	//根据反射的值，来获取对应的类型和数值
	fmt.Println("--------------")
	v := reflect.ValueOf(x)
	fmt.Println("kind is float:", v.Kind() == reflect.Float64)
	fmt.Println("type:", v.Type())
	fmt.Println("value:", v.Float())
}
