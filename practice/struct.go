package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Circal struct {
	x float64
	y float64
	r float64
}

func (c *Circal) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circal) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func totalArea(Shapes ...Shape) float64 {
	var area float64
	for _, s := range Shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	c := Circal{0, 0, 10}
	fmt.Println(c.area())
	fmt.Println(c.perimeter())

	fmt.Println(totalArea(&r, &c))

	m := MultiShape{[]Shape{&r, &c}}
	fmt.Println(m.area())
}
