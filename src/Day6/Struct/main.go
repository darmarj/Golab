package main

import "fmt"

type Book struct {
	id      int
	title   string
	author  string
	subject string
}

func printBook(book *Book) {
	fmt.Printf("id=%d, title=%s, author=%s, subject=%s\n",
		book.id, book.title, book.author, book.subject)
	book.id = 1000
}

func main() {
	var book1 *Book
	book1 = new(Book)
	book1.id = 1001
	book1.title = "go in action"
	book1.author = "John"
	book1.subject = "about GOlang"
	//fmt.Println(book1)

	book2 := Book{
		id:      1002,
		title:   "python in action",
		author:  "Ken",
		subject: "about Python",
	}
	//fmt.Println(book2)
	//fmt.Println("book2.title=", book2.title)

	printBook(&book2)
	fmt.Println(book2)

}
