package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
	Color        string
	X, Y, Radius int
}

func (c *Circle) SetX(x int) {
	c.X = x
}

func (c *Circle) SetY(y int) {
	c.Y = y
}

func (c *Circle) SetRadius(radius int) {
	c.Radius = radius
}

func (c *Circle) Draw() {
	fmt.Println("Draw circle, color: ", c.Color, ", x, y, radius: ", c.X, c.Y, c.Radius)
}

type ShapeFactory struct {
	Hash map[string]Shape
}

func (s *ShapeFactory) GetCircle(color string) *Circle {
	if v, ok := s.Hash[color]; ok {
		return v.(*Circle)
	}
	circle := &Circle{Color: color}
	s.Hash[color] = circle
	fmt.Println("Creating circle: ", circle)
	return circle
}

func main() {
	colors := []string{"red", "green", "blue", "white", "black"}
	factory := &ShapeFactory{Hash: map[string]Shape{}}
	for i := 0; i < 5; i++ {
		factory.GetCircle(colors[i])
	}

	for _, v := range factory.Hash {
		fmt.Println("v: ", v)
	}
}
