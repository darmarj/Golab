package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
	Sex string
}
func (p Person) Say(msg string) {
	fmt.Printf("Hey", msg)
}
func (p Person) PrintInfo(msg string) {
	fmt.Printf("Name: %s, Age: %d, Sex: %s\n", p.Name, p.Age, p.Sex)
}

func main() {
	p1 := Person{ Name: "John", Age: 30, Sex: "Male"}
	GetMessage(p1)
}
// 获取input的信息
func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input) //先获取input的类型
	fmt.Println("get Type is :", getType.Name()) //Person
	fmt.Println("get Kind is :", getType.Kind()) //struct

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is :", getValue)

	// 获取字段
	/*
	step1: 先获取Type对象: reflect.Type,
		NumField()
		Field(index)
	Step2: 通过Field()获取每一个Field字段
	Step3: Interface()，得到对应的Value
	*/
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface() //获取第一个数据
		fmt.Printf("字段名称: %s, 字段类型: %s, 字段数值: %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称: %s, 方法类型: %v\n", method.Name, method.Type)
	}
}

