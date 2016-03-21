package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, s := range shapes {
		area += s.area()

	}

	return area
}

func totalPerimeter(shapes ...Shape) float64 {

	var perimeter float64

	for _, s := range shapes {

		perimeter += s.perimeter()
	}

	return perimeter
}

func (r *Rectangle) area() float64 {

	l := distance(r.x1, r.x1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) perimeter() float64 {

	l := distance(r.x1, r.x1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return 2 * (l + w)

}

func (c *Circle) perimeter() float64 {

	return 2 * math.Pi * c.r

}

func main() {

	//var rx1, ry1 float64 = 0, 0
	//var rx2, ry2 float64 = 10, 10
	//var cx, cy, cr float64 = 0, 0, 5
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}

	fmt.Println(r.area(), r.perimeter())
	fmt.Println(c.area(), c.perimeter())

	fmt.Println(totalArea(&c, &r))
	fmt.Println(totalPerimeter(&c, &r))

}
