package structsmethodsinterfaces

import "math"

type (
	Rectangle struct {
		Width  float64
		Height float64
	}

	Circle struct {
		Radius float64
	}

	Triangle struct {
		Base   float64
		Height float64
	}
)

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * math.Pi * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
