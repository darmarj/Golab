package main

import (
	"fmt"
	"math"
	// "github.com/xyz/lib1/T1"
)

// package: github.com/xyz/lib2/T1
// type T1 = lib2.T1

//MAX is working for export
const (
	s1  = 100 //Private Export:只能在当前目录下使用的常量
	MAX = 999 //Public Export:其他目录可以引用的常量，字母大写
)

func enumDemo1() {
	const (
		// Sunday    = 0
		// Monday    = 1
		// Tuesday   = 2
		// Wednesday = 3
		// Thursday  = 4
		// Friday    = 5
		// Saturday  = 6
		Sunday = 1 + iota // 1-->2-->3...
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func enumDemo2() {
	const (
		// Sunday    = 0
		// Monday    = 1
		// Tuesday   = 2
		// Wednesday = 3
		// Thursday  = 4
		// Friday    = 5
		// Saturday  = 6
		Monday = 1 + iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}

func typeDefAndAlias() {
	type MyInt1 int   //Define only for documentation named as "MyInt1", usually for coding (alias)
	type MyInt2 = int //MyInt2 is the same as int

	var i int = 25
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func constDemo() {
	const s string = "YES"
	const a, b = 3, 4
	fmt.Println(s, a, b)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)

	fmt.Println(s1, MAX)

}

//定义一组常量
const C1, C2, C3 = 100, 3.14, "LOL"
const (
	MALE   = 0
	FEMALE = 1
	UNKOWN = 2
)

//一组常量中，如果某个常量没有初始值，默认和上一行一致
func groupDemo() {
	const (
		a int = 100
		b
		c string = "Good"
		d
		e
	)
	fmt.Println("----------------------")
	fmt.Printf("%T, %d\n", a, a)
	fmt.Printf("%T, %d\n", b, b)
	fmt.Printf("%T, %s\n", c, c)
	fmt.Printf("%T, %s\n", d, d)
	fmt.Printf("%T, %s\n", e, e)
}

// 枚举类型: 使用常量作为类型，一组相关数值的数据
func enumDemo3() {
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 4
	)
	fmt.Println("SPRING, SUMMER, AUTUMN, WINTER")
}

func main() {
	constDemo()
	enumDemo1()
	enumDemo2()
	typeDefAndAlias()
	groupDemo()
	enumDemo3()
}