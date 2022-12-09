package main

import "fmt"

type humanoid interface {
	speak()
	walk()
}

type human struct {
	name string
}

func (h human) speak() {
	fmt.Println(h.name, " speak")
}

func (h human) walk() {
	fmt.Println(h.name, " walk")
}

func (h human) String() string {
	return "hello, " + h.name
}

type dog struct {
	alias string
}

func (d dog) walk() {
	fmt.Println(d.alias, " walk")
}

func mainDefineInterface() {
	d := dog{"kuki"}

	h := human{"Alex"}
	fmt.Println(h)

	d.walk()
	h.walk()

	doHumanoid(h)
	//doHumanoid(d)
}

func doHumanoid(h humanoid) {
	h.speak()
	h.walk()
}
