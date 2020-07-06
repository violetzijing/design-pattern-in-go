package main

import (
	"fmt"
)

type ShapeWithColor struct {
}

type Shape interface {
	Draw()
}

type Color interface {
	Fill()
}

type Rectangle struct{}
type Square struct{}
type Circle struct{}

type Red struct{}
type Green struct{}
type Blue struct{}

func (r *Rectangle) Draw() {
	fmt.Println("rectangle draw")
}

func (s *Square) Draw() {
	fmt.Println("square draw")
}

func (c *Circle) Draw() {
	fmt.Println("circle draw")
}

func (r *Red) Fill() {
	fmt.Println("fill in red")
}

func (g *Green) Fill() {
	fmt.Println("fill in green")
}

func (b *Blue) Fill() {
	fmt.Println("fill in blue")
}

func (s *ShapeWithColor) GetColor(color string) Color {
	switch color {
	case "red":
		return &Red{}
	case "green":
		return &Green{}
	case "blue":
		return &Blue{}
	}
	return nil
}

func (s *ShapeWithColor) GetShape(str string) Shape {
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
	instance := &ShapeWithColor{}

	shape1 := instance.GetShape("circle")
	shape1.Draw()

	color1 := instance.GetColor("red")
	color1.Fill()

	shape2 := instance.GetShape("square")
	shape2.Draw()

	color2 := instance.GetColor("green")
	color2.Fill()

	shape3 := instance.GetShape("rectangle")
	shape3.Draw()

	color3 := instance.GetColor("blue")
	color3.Fill()
}
