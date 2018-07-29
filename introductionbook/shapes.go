package main

import (
	"math"
	"fmt"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi*c.r*c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	length := distance(r.x1, r.y1, r.x2, r.y1)
	width := distance(r.x1, r.y1, r.x1, r.y2)

	return length*width
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func totalArea(shapes ...Shape) float64 {
	total := float64(0)
	for _, s := range shapes {
		total += s.area()
	}

	return total
}

func main() {
	var c *Circle
	c = &Circle{0, 0, 5}
	fmt.Printf("Area of r=%f is %f\n", c.r, c.area())

	r := &Rectangle{0,0, 10, 20}
	fmt.Printf("Area of rectangle is %f\n", r.area())

	fmt.Println("Total area is:", totalArea(c, r))
}