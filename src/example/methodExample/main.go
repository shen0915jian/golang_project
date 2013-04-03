package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	with, height float64
}

type Circle struct {
	redius float64
}

func (r Rectangle) area() float64 {
	return r.with * r.height
}

func (c Circle) area() float64 {
	return c.redius * c.redius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is ", r1.area())
	fmt.Println("Area of r2 is :", r2.area())
	fmt.Println("Area of c1 is :", c1.area())
	fmt.Println("Area of c2 is :", c2.area())

}
