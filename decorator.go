package main

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct{}

func (*Rectangle) Draw() {
	fmt.Println("draw rectangle")
}

type Circle struct{}

func (*Circle) Draw() {
	fmt.Println("draw circle")
}

type ShapeDecorator struct {
	Shape Shape
}

func NewShapeDecorator(shape Shape) *ShapeDecorator {
	return &ShapeDecorator{
		Shape: shape,
	}
}

func (s *ShapeDecorator) Draw() {
	s.Shape.Draw()
}

type RedShapeDecorator struct {
	Shape *ShapeDecorator
}

func NewRedShapeDecorator(shape Shape) *RedShapeDecorator {
	return &RedShapeDecorator{
		Shape: &ShapeDecorator{Shape: shape},
	}
}

func (*RedShapeDecorator) SetBorderColor() {
	fmt.Println("Border color: red")
}

func (r *RedShapeDecorator) Draw() {
	r.Shape.Draw()
	r.SetBorderColor()
}

func main() {
	circle := &Circle{}
	redCircle := NewRedShapeDecorator(&Circle{})
	redRectangle := NewRedShapeDecorator(&Rectangle{})

	fmt.Println("===")
	circle.Draw()
	fmt.Println("===")
	redCircle.Draw()
	fmt.Println("===")
	redRectangle.Draw()
}
