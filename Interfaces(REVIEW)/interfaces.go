package main

import (
	"fmt"
	"math"
)

type goemetry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.width
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g goemetry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
	fmt.Println("-----")
}

func detectCircle(g goemetry) {
	c, ok := g.(circle)
	fmt.Println("HERE: ", c, g, ok)
	if ok {
		fmt.Println("circle with radius", c.radius)
	} else {
		fmt.Println("rect with perim", g.perim())
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
