package main

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct{}
type Square struct{}
type Circle struct{}

func (*Rectangle) Draw() {
	fmt.Println("draw rectangle")
}

func (*Square) Draw() {
	fmt.Println("draw square")
}

func (*Circle) Draw() {
	fmt.Println("draw circle")
}

type ShapeMaker struct {
	circle    *Circle
	rectangle *Rectangle
	square    *Square
}

func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{
		circle:    &Circle{},
		rectangle: &Rectangle{},
		square:    &Square{},
	}
}

func (s *ShapeMaker) DrawCircle() {
	s.circle.Draw()
}

func (s *ShapeMaker) DrawRectangle() {
	s.rectangle.Draw()
}

func (s *ShapeMaker) DrawSquare() {
	s.square.Draw()
}

func main() {
	maker := NewShapeMaker()
	maker.DrawCircle()
	maker.DrawRectangle()
	maker.DrawSquare()
}
