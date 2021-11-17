package main

import "fmt"

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func Counter(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func IsLarge(l Lener) bool {
	return l.Len() > 50
}

func main() {
	var lst List
	Counter(&lst, 1, 10)
	fmt.Println(lst)
	//
	//point
	plst := new(List)
	Counter(plst, 1, 10)
	fmt.Println(plst)

	IsLarge(plst)
}
