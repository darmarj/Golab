package main

import "fmt"

func changeElemnet(arr []int) {
	arr[0] = 1000
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s01 := arr[2:7]
	fmt.Println(s01)

	s := arr[3]
	fmt.Println(s)

	s02 := arr[:6]
	fmt.Println(s02)

	s03 := arr[2:]
	fmt.Println(s03)

	s04 := arr[:]
	fmt.Println(s04)

	s05 := make([]int, 5, 10)
	fmt.Printf("s5=%v len=%d cap=%d\n", s05, len(s05), cap(s05))

	s06 := arr[2:6]
	fmt.Printf("arr=%v s06=%v len=%d cap=%d\n", arr, s06, len(s06), cap(s06))

	s07 := s06[3:8]
	fmt.Printf("arr=%v s07=%v len=%d cap=%d\n", arr, s07, len(s07), cap(s07))

	changeElemnet(s07)
	fmt.Printf("s07=%v\ns06=%v\narr=%v\n", s07, s06, arr)
}
