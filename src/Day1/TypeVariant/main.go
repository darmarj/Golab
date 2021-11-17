package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

func main() {
	var a int = 1
	var b int32 = 2
	// b = int32(a)
	fmt.Printf("%T %T\n", a, b)
	fmt.Println(a, b)

	cpuArch := runtime.GOARCH
	intSize := strconv.IntSize
	fmt.Println(cpuArch, intSize)

	var f1 float32
	var f2 float64
	fmt.Printf("%F %F\n", f1, f2)

	var bt byte = 2
	var ru rune = '好'
	fmt.Printf("%T %T\n", bt, ru)

	var a1, a2 = 3, 4
	var c int
	temp := a1*a2 + a2*a2
	fmt.Printf("%T\n", temp)
	c = int(math.Sqrt(float64(temp)))
	fmt.Printf("%T %d\n", c, c)

	p := 1
	ptr := &p
	fmt.Printf("%T %T", p, ptr)
}
