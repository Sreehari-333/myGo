package main

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (c Circle) Area() float64 {

	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {

	return 2 * 3.14 * c.radius
}

func (r Rectangle) Area() float64 {

	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {

	return 2 * (r.length + r.breadth)
}

func main() {

	circle1 := Circle{
		radius: 5,
	}

	rectangle1 := Rectangle{
		length:  10,
		breadth: 5,
	}

	fmt.Printf("Area of circle is %.1f\n", circle1.Area())
	fmt.Printf("Perimeter of the circle is %.1f\n", circle1.Perimeter())

	fmt.Printf("Area of Rectangle is %.1f\n", rectangle1.Area())
	fmt.Printf("Perimeter of Rectangle is %.1f\n", rectangle1.Perimeter())

}
