package main

import "fmt"

type ComputerPart interface {
	Accept(ComputerPartVisitor)
}

type ComputerPartVisitor interface {
	Visit(ComputerPart)
}

type Keyboard struct{}

func (k *Keyboard) Accept(visitor ComputerPartVisitor) {
	visitor.Visit(k)
}

type Monitor struct{}

func (m *Monitor) Accept(visitor ComputerPartVisitor) {
	visitor.Visit(m)
}

type Mouse struct{}

func (m *Mouse) Accept(visitor ComputerPartVisitor) {
	visitor.Visit(m)
}

type Computer struct {
	ComputerParts []ComputerPart
}

func NewComputer() *Computer {
	return &Computer{
		ComputerParts: []ComputerPart{&Keyboard{}, &Monitor{}, &Mouse{}},
	}
}

func (c *Computer) Accept(visitor ComputerPartVisitor) {
	for _, p := range c.ComputerParts {
		p.Accept(visitor)
	}
	visitor.Visit(c)
}

type ComputerPartDisplayVisitor struct{}

func (c *ComputerPartDisplayVisitor) Visit(part ComputerPart) {
	switch part.(type) {
	case *Computer:
		fmt.Println("Displaying computer")
	case *Keyboard:
		fmt.Println("Displaying keyboard")
	case *Monitor:
		fmt.Println("Displaying monitor")
	case *Mouse:
		fmt.Println("Displaying mouse")
	}
}

func main() {
	computer := NewComputer()
	computer.Accept(&ComputerPartDisplayVisitor{})
}
