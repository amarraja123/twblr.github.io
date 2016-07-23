package interfaces

import "fmt"

type Shape interface {
	Area() int
}

type Square struct {
	side int
}

type Rectangle struct {
	length, breadth int
}

type Hybrid struct {
	s Shape
	r Shape
}

func (s Square) Area() int {
	return findType(s)
}

func (r Rectangle) Area() int {
	return findType(r)
}

func (h Hybrid) Area() int {
	return findType(h)
}

func findType(h Shape) int {

	switch t := h.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
		return 0
	case Square:
		return t.side*t.side
	case Rectangle:
		return t.length*t.breadth
	case Hybrid:
		return t.s.Area() + t.r.Area()

	}
}

