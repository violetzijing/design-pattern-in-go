package main

import "fmt"

type State interface {
	DoAction(*Context)
	ToString() string
}

type StartState struct{}

func (s *StartState) DoAction(ctx *Context) {
	fmt.Println("Player is in start state")
	ctx.SetState(s)
}

func (s *StartState) ToString() string {
	return "start state"
}

type EndState struct{}

func (e *EndState) DoAction(ctx *Context) {
	fmt.Println("Player is in end state")
	ctx.SetState(e)
}

func (e *EndState) ToString() string {
	return "stop state"
}

type Context struct {
	State State
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetState(state State) {
	c.State = state
}

func (c *Context) GetState() State {
	return c.State
}

func main() {
	ctx := NewContext()

	startState := &StartState{}
	startState.DoAction(ctx)

	fmt.Println(ctx.GetState().ToString())

	stopState := &EndState{}
	stopState.DoAction(ctx)

	fmt.Println(ctx.GetState().ToString())
}
