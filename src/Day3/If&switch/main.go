package main

import "fmt"

func main() {
	a := 26
	b := 2

	if a < 20 {
		fmt.Println("a < 20")
	} else if a == 20 {
		fmt.Println("a >= 20")
	} else {
		fmt.Println("a > 20")
		if b > 5 {
			fmt.Println("b > 5")
		} else {
			fmt.Println("b <= 5")
		}
	}

	fmt.Println("--------------")

	switch a {
	case 20:
		fmt.Println("a = 20")
	case 21:
		fmt.Println("a = 21")
	case 22:
		fmt.Println("a = 22")
	case 25, 26, 27:
		fmt.Println("a in [25, 26, 27]")
	default:
		fmt.Println("Unkown")
	}

	switch {
	case a == 20:
		fmt.Println("a = 20")
	case a > 20:
		fmt.Println("a > 20")
		fallthrough // 不管之后的判断是否是真或假, 直接运行下一步命令
	case a < 20:
		fmt.Println("a < 20")
	default:
		fmt.Println("Unkown")
	}
}
