package main

import "fmt"

type Strategy interface {
	DoOperation(int, int) int
}

type OperationAdd struct{}

func (o *OperationAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}

type OperationSubtract struct{}

func (o *OperationSubtract) DoOperation(num1, num2 int) int {
	return num1 - num2
}

type OperationMultiply struct{}

func (o *OperationMultiply) DoOperation(num1, num2 int) int {
	return num1 * num2
}

type Context struct {
	Strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{
		Strategy: strategy,
	}
}

func (c *Context) Execute(num1, num2 int) int {
	return c.Strategy.DoOperation(num1, num2)
}

func main() {
	ctx := NewContext(&OperationAdd{})
	fmt.Println("10 + 5 = ", ctx.Execute(10, 5))

	ctx = NewContext(&OperationSubtract{})
	fmt.Println("10 - 5 = ", ctx.Execute(10, 5))

	ctx = NewContext(&OperationMultiply{})
	fmt.Println("10 * 5 = ", ctx.Execute(10, 5))
}
