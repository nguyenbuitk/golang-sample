package main

import (
	"fmt"
	"math"
)

// Golang không phải ngôn ngữ OOP, không có class, inheritance, ...
// Nhưng tương tự như kế thừa trong OOP, circle và rectangle được kế thừa từ Shape
// Có thể hình dung đây là class Shape
type Shape interface {
	Area() float64
}

// Có thể hình dung đây là class Circle
type Circle struct {
	Radius float64
}

// function này có cùng tên (Area()) và kiểu trả về (float64) đối với interface shape, nên nó là một member của shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var s Shape
	c := Circle{Radius: 2.5}
	fmt.Println(c.Area())

	s = Circle{Radius: 2.5}
	fmt.Println(s.Area())

	s = Rectangle{Width: 2.0, Height: 3.0}
	fmt.Println(s.Area())
}
