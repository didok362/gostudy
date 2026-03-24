package main

import (
	"fmt"
	"math"
)

type figure interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return 4 * s.side
}

type Rectangle struct {
	side1 float64
	side2 float64
}

func (r Rectangle) area() float64 {
	return r.side1 * r.side2
}

func (r Rectangle) perimeter() float64 {
	return 2*r.side1 + 2*r.side2
}

func info(f figure) {
	fmt.Printf("Площадь: %f\n", f.area())
	fmt.Printf("Периметр: %f\n", f.perimeter())
}
func main() {
	Krug1 := Circle{Radius: 5}
	Kvadrat1 := Square{side: 12}
	Pryamougolnic := Rectangle{side1: 12, side2: 24}
	info(Krug1)
	info(Kvadrat1)
	info(Pryamougolnic)
}
