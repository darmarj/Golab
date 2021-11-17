package main

import "fmt"

func changeElemnet(arr *[3]int) {
	arr[0] = 1000
}

func main() {
	var arr01 [5]int
	arr01[1] = 10
	fmt.Println(arr01, arr01[0], arr01[1])

	arr02 := [3]int{1, 2, 3}
	fmt.Println(arr02)

	arr03 := [...]int{5, 6, 7, 8, 9, 12}
	fmt.Println(arr03)
	fmt.Println(arr03, arr03[4])
	fmt.Println(len(arr03), arr03)
	fmt.Println("==============")
	for i := 0; i <= len(arr03)-1; i++ {
		fmt.Println(arr03[i])
	}
	fmt.Println("==============")
	for i, num := range arr03 {
		fmt.Println(i, num)
	}
	fmt.Println("==============")
	for _, num := range arr03 {
		fmt.Println(num)
	}
	fmt.Println("==============")
	for i := range arr03 {
		fmt.Println(i)
	}
	fmt.Println("==============")
	var grid = [5][6]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(grid)
	fmt.Println("==============")

	arr05 := [3]int{1, 2, 3}
	changeElemnet(&arr05)
	fmt.Println(arr05)
}
