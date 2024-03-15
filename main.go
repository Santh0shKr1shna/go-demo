package shape

import (
	"errors"
)

type Shape interface {
	CalculateArea()
	CalculatePerimeter()
}

type Rectangle struct {
	height int
	breadth int
}

func NewRectangle(height int, breadth int) (*Rectangle, error) {
	if height > 0 && breadth > 0 {
		r := Rectangle{
			height: height,
			breadth: breadth,
		}
		return &r, nil
	}
	return nil, errors.New("Rectangle cannot have negative or zero dimensions")
}

func (r Rectangle) GetHeight() int {
	return r.height
}

func (r Rectangle) GetBreadth() int {
	return r.breadth
}

func (r Rectangle) CalculateArea() int {
	return r.breadth * r.height
}

func (r Rectangle) CalculatePerimeter() int {
	return 2 * (r.breadth + r.height)
}

// SQUARE

type Square struct {
	side int
}

func NewSquare (side int) (*Square, error) {
	if side > 0 {
		sq := Square {
			side: side,
		}
		return &sq, nil
	}
	return nil, errors.New("Square cannot have negative or zero sides")
}

func (s Square) CalculateArea() int {
	return s.side * s.side
}

func (s Square) CalculatePerimeter() int {
	return 4 * s.side
}

func (s Square) GetSide() int {
	return s.side
}