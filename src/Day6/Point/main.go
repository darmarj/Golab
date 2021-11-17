package main

import (
	"fmt"
)

// var outside
var g int = 100

// a, b is
func sum(a, b int) int {
	// var c is inside
	var c = 10
	s := a + b + c
	return s
}

func funcValRef(a int) int {
	a = 1000
	return a
}

func funcValRefPrt(a *int) {
	*a = 1000
	fmt.Printf("inside is: %d\n", *a)
}

func main() {
	// var inside

	println(int(sum(5, 6)))
	fmt.Println("==============")

	var a, b, c int
	g = 200

	a = 10
	b = 20
	c = a + b + g
	println(c)
	fmt.Println("==============")

	//var p int = 100
	//funcValRef(p)
	//fmt.Printf("funvalref inside is: %d\n", funcValRef(a))
	//fmt.Println(p)
	//fmt.Println("==============")

	var p int = 100
	fmt.Println(p)
	funcValRefPrt(&p)
	//fmt.Println(p)

}
