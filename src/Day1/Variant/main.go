/*
	Tips:
	* 变量必须先定义才能使用
	* 变量的类型和赋值必须一致
	* 同一个作用域内，变量不能冲突
	* 简短定义方式，不能定义全局变量
	* 简短定义方式，左边的变量至少有一个新的
	* 变量的零值，就是默认值
		整型: 默认值是 0
		浮点类型: 默认值是 0
		字符串: 默认值是 ""，空
*/
package main

import "fmt"

var c1, c2 = 6, "gogogo"

// d1, d2 := 7, "cool..."

var (
	d1 = 7
	d2 = "cool..."
)

func varInitValue() {
	var a, b int = 3, 4
	fmt.Println(a, b)
	var s string = "abc"
	fmt.Println(s)

	var s1, s2, s3 = 5, "hello", true
	fmt.Println(s1, s2, s3)

	s4, s5, s6 := 6, "world", false
	fmt.Println(s4, s5, s6)
}

func main() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
	// var a int = 3
	// a = 3
	varInitValue()
	fmt.Println(c1, c2)
	fmt.Println(d1, d2)

	num := 100
	fmt.Printf("num is :%d, address is :%p\n", num, &num)
	num = 200 //no new variables on left side of :=
	fmt.Printf("num is :%d, address is :%p\n", num, &num)

}
