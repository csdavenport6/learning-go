package main

import "fmt"

type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
}

// A Var identifies a variable, e.g. x
type Var string

// A literal is a numeric constant, e.g. 3.141.
type literal float64

// A unary represents a unary operator expression, e.g. -x
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression, e.g. x+y
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g. sin(x)
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func main() {
	fmt.Println("vim-go")
}