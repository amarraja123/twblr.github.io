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
	s Square
	r Rectangle
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
		square := Square(t)
		return square.side*square.side
	case Rectangle:
		rectangle := Rectangle(t)
		return rectangle.length*rectangle.breadth
	case Hybrid:
		h := Hybrid(t)
		return (h.s.side * h.s.side) + (h.r.breadth * h.r.length)

	}
}

