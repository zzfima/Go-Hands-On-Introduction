package main

import "fmt"

type author struct {
	name string
}

type book struct {
	title  string
	author author
}

type library struct {
	books []book
}

func (lib library) addBook(newBook book) {
	lib.books = append(lib.books, newBook)
}

func (lib library) lookupByAuthorName(authorName string) (authorBooks []book) {
	for _, b := range lib.books {
		if b.author.name == authorName {
			authorBooks = append(authorBooks, b)
		}
	}

	return
}

func mainCompositeTypes() {
	a1 := author{name: "Alex"}
	a2 := author{name: "Sofy"}

	b1 := book{"Algebra1", a1}
	b2 := book{"Algebra2", a1}
	b3 := book{"Algebra3", a1}
	b4 := book{"Algebra4", a1}

	b5 := book{"Algo1", a2}
	b6 := book{"Algo2", a2}
	lib := library{[]book{b1, b2, b3, b4, b5, b6}}
	b7 := book{"Algo3", a2}
	lib.addBook(b7)

	fmt.Println(lib.lookupByAuthorName("Alex"))
	fmt.Println(lib.lookupByAuthorName("Sofy"))
}
