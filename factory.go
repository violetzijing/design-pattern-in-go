package main

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type Rectangle struct{}

type Square struct{}

type Circle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("rectangle draw")
}

func (s *Square) Draw() {
	fmt.Println("square draw")
}

func (c *Circle) Draw() {
	fmt.Println("circle draw")
}

func GetShape(str string) Shape {
	switch str {
	case "circle":
		return &Circle{}
	case "rectangle":
		return &Rectangle{}
	case "square":
		return &Square{}
	}

	return nil
}

func main() {
	shape1 := GetShape("rectangle")
	shape1.Draw()

	shape2 := GetShape("square")
	shape2.Draw()

	shape3 := GetShape("circle")
	shape3.Draw()
}
