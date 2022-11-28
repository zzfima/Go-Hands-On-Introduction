package main

import "fmt"

//Shape describes shape.
type Shape struct {
	name, description string
	size              int
}

func main() {
	s1 := new(Shape)
	PrintShape(s1)

	s2 := Shape{
		"rect", "shape kind", 34,
	}
	PrintShape(&s2)

	s3 := Shape{
		description: "shape kind", name: "rect", size: 33,
	}
	PrintShape(&s3)

	s4 := NewShape()
	PrintShape(&s4)
}

//NewShape shape factory
func NewShape() Shape {
	return Shape{
		"no name", "no desc", 0,
	}
}

//PrintShape Shape ToString
func PrintShape(s *Shape) {
	fmt.Printf("name: %s, description: %s, size: %d\n", s.name, s.description, s.size)
}
