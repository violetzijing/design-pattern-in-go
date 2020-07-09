package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Interpreter(context string) bool
}

type TerminalExpression struct {
	Data string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{
		Data: data,
	}
}

func (t *TerminalExpression) Interpreter(context string) bool {
	return strings.Contains(context, t.Data)
}

type OrExpression struct {
	Exp1 Expression
	Exp2 Expression
}

func NewOrExpression(exp1, exp2 Expression) *OrExpression {
	return &OrExpression{
		Exp1: exp1,
		Exp2: exp2,
	}
}

func (o *OrExpression) Interpreter(context string) bool {
	return o.Exp1.Interpreter(context) || o.Exp2.Interpreter(context)
}

type AndExpression struct {
	Exp1 Expression
	Exp2 Expression
}

func NewAndExpression(exp1, exp2 Expression) *AndExpression {
	return &AndExpression{
		Exp1: exp1,
		Exp2: exp2,
	}
}

func (a *AndExpression) Interpreter(context string) bool {
	return a.Exp1.Interpreter(context) && a.Exp2.Interpreter(context)
}

func main() {
	robert := NewTerminalExpression("Robert")
	john := NewTerminalExpression("John")
	orExpression := NewOrExpression(robert, john)

	julie := NewTerminalExpression("Julie")
	married := NewTerminalExpression("married")
	andExpression := NewAndExpression(julie, married)

	isMale := orExpression
	isMarriedWoman := andExpression
	fmt.Println("Is John male? ", isMale.Interpreter("John"))
	fmt.Println("Julie is married woman? ", isMarriedWoman.Interpreter("married Julie"))
}
