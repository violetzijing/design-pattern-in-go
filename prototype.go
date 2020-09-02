package main

import "fmt"

type Shape struct {
	ID   int
	Type string
}

func (s *Shape) Draw() string {
	return s.Type
}

func (s *Shape) GetID() int {
	return s.ID
}

func (s *Shape) SetID(ID int) {
	s.ID = ID
}

func Clone(input *Shape) *Shape {
	return &Shape{
		ID:   input.ID,
		Type: input.Type,
	}
}

type Rectangle struct {
	Shape *Shape
}

func NewRectangle() *Rectangle {
	return &Rectangle{
		Shape: &Shape{
			Type: "Rectangle",
		},
	}
}

func (r *Rectangle) Draw() {
	fmt.Println("draw Rectangle")
	fmt.Println("id: ", r.Shape.ID)
}

type Square struct {
	Shape *Shape
}

func NewSquare() *Square {
	return &Square{
		Shape: &Shape{
			Type: "Square",
		},
	}
}

func (s *Square) Draw() {
	fmt.Println("draw Square")
	fmt.Println("id: ", s.Shape.ID)
}

type Circle struct {
	Shape *Shape
}

func NewCircle() *Circle {
	return &Circle{
		Shape: &Shape{
			Type: "Circle",
		},
	}
}

func (c *Circle) Draw() {
	fmt.Println("draw Circle")
	fmt.Println("id: ", c.Shape.ID)
}

func LoadCache() map[string]interface{} {
	hash := map[string]interface{}{}
	circle := NewCircle()
	circle.Shape.SetID(1)
	hash["circle"] = circle

	rectangle := NewRectangle()
	rectangle.Shape.SetID(2)
	hash["rectangle"] = rectangle

	square := NewSquare()
	square.Shape.SetID(3)
	hash["square"] = square

	return hash
}

func GetShape(shape interface{}) interface{} {
	switch v := shape.(type) {
	case *Circle:
		return &Circle{Shape: Clone(v.Shape)}
	case *Rectangle:
		return &Rectangle{Shape: Clone(v.Shape)}
	case *Square:
		return &Square{Shape: Clone(v.Shape)}
	}
	return nil
}

func main() {
	cache := LoadCache()
	GetShape(cache["circle"]).(*Circle).Draw()
	GetShape(cache["rectangle"]).(*Rectangle).Draw()
	GetShape(cache["square"]).(*Square).Draw()
}
