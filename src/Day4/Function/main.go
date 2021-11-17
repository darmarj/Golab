package main

import (
	"fmt"
	"os"
)

func operator(a, b int, op string) int {
	// +-*/
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic(fmt.Sprintf("not support operate: %s", op))
	}
}

func swap(a, b int) (x int, y int) {
	x, y = b, a
	return
}

func main() {
	a, b := 10, 5
	fmt.Println(operator(a, b, "+"))
	fmt.Println(operator(a, b, "-"))
	fmt.Println(operator(a, b, "*"))
	fmt.Println(operator(a, b, "/"))
	//fmt.Println(operator(a, b, "@"))
	x, y := swap(a, b)
	fmt.Println(x, y)
	fmt.Println(swap(a, b))

	if file, err := os.Open("abc.txt"); err != nil {
		fmt.Println("打开文件有误")
	} else {
		fmt.Println(file)
	}
}
