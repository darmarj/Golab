package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	a := 15
	b := 10
	for a >= b { //while circulation feature
		b++
		fmt.Println(b)
	}
	/*
		for {
			fmt.Println("dead loop")
		}

		fmt.Println("testing") // dead loop is always in looping, never run "testing"
	*/

	for a < 20 {
		fmt.Printf("a =%d\n", a)
		a++
		if a > 17 {
			break // jump out to the loop
		}
	}

	fmt.Println("-------------")

	a = 15
	for a < 20 {
		a++
		if a == 17 {
			continue //jump out to the active loop
		}
		fmt.Printf("a =%d\n", a)
	}

	fmt.Println("-------------")

	a = 10
LOOP:
	for a < 20 {
		a++
		if a == 15 {
			goto LOOP
		}
		fmt.Printf("a =%d\n", a)
	}
}
