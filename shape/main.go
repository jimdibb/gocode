package main

import "fmt"
import "math"

type shape interface {
	Perimeter() float32
	Area() float32
	Sides() float32
}

type square struct {
	length float32
}

type factory interface {
	Make(t string, s float32) shape
}

type shapefactory struct{}

func (f shapefactory) Make(t string, size float32) shape {
	if t == "square" {
		return square{length: size}
	}
	if t == "circle" {
		return circle{radius: size}
	}
	return nil
}

func (s square) Sides() float32 {
	return float32(4)
}
func (s square) Perimeter() float32 {
	return s.Sides() * s.length
}
func (s square) Area() float32 {
	return s.length * s.length
}

type circle struct {
	radius float32
}

func (c circle) Sides() float32 {
	return float32(1)
}
func (c circle) Perimeter() float32 {
	return math.Pi * 2 * c.radius
}
func (c circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var f shapefactory
	s := f.Make("square", 5)
	c := f.Make("circle", 1)

	fmt.Println("area", s.Area())
	fmt.Println("Perimeter", s.Perimeter())
	fmt.Println("area", c.Area())
	fmt.Println("Perimeter", c.Perimeter())

}
