package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}

func (u User) String() {
	fmt.Printf("Name= %s, Age= %d\n", u.Name, u.Age)
}

func (u User) Message() {
	fmt.Println("Hey..")
}

func reflectBasic() {
	u := User{"hyden", 30}

	t0 := reflect.TypeOf(u)
	fmt.Printf("Typeof(u)= %v\n", t0)

	v := reflect.ValueOf(u)
	t1 := v.Type()
	fmt.Printf("refelct.Value(u)= %v, reflect.Type(u)= %v\n", v, t1)

	fmt.Printf("%T, %v\n", u, u)

	fmt.Println("--------------------")
	// reflect.Value 转换成原始数据
	u1 := v.Interface().(User)
	fmt.Println(u1, reflect.TypeOf(u1))

	fmt.Println("--------------------")
	// 获取底层的数据型
	fmt.Println(v.Type().Kind())
}

func reflectLoopStruct() {
	u := User{"Mike", 20}

	t := reflect.TypeOf(u)
	for i :=0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("fieldName: %s\n", f.Name)
		fmt.Printf("fieldIndex: %d, fieldName: %s\n", f.Index, f.Name)
	}

	for i :=0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("methodName: %s\n", m.Name)
	}
}

func main() {
	//reflectBasic()
	reflectLoopStruct()
}