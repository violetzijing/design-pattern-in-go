package main

import "fmt"

type DrawAPI interface {
	DrawCircle(int, int, int)
}

type RedCircle struct{}
type GreenCircle struct{}

func (*RedCircle) DrawCircle(radius, x, y int) {
	fmt.Println("draw red circle, radius, x, y: ", radius, x, y)
}

func (*GreenCircle) DrawCircle(radius, x, y int) {
	fmt.Println("draw green circle, radius, x, y: ", radius, x, y)
}

type Shape struct {
	drawAPI DrawAPI
}

func (s *Shape) Draw(x, y, radius int) {
	s.drawAPI.DrawCircle(x, y, radius)
}

type Circle struct {
	x, y, radius int
	shape        *Shape
}

func NewCircle(x, y, radius int) *Circle {
	return &Circle{
		x:      x,
		y:      y,
		radius: radius,
	}
}

func (c *Circle) Draw() {
	c.shape.Draw(c.x, c.y, c.radius)
}

func main() {
	circle := NewCircle(1, 2, 3)
	circle.shape = &Shape{drawAPI: &RedCircle{}}
	circle.Draw()

	circle.shape = &Shape{drawAPI: &GreenCircle{}}
	circle.Draw()
}
