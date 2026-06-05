package main

import "math"

func (n Number) Eval() float64 {
	return n.Value
}

func (b BinaryExpr) Eval() float64 {
	switch b.op {
	case "+":
		return b.Left.Eval() + b.Right.Eval()
	case "-":
		return b.Left.Eval() - b.Right.Eval()
	case "*":
		return b.Left.Eval() * b.Right.Eval()
	case "/":
		return b.Left.Eval() / b.Right.Eval()
	case "^":
		return math.Pow(
			b.Left.Eval(),
			b.Right.Eval(),
		) 
	}
	panic("unknown operator")
}

func factorial(inp float64) float64 {
	minFlag := false
	if inp < 0 {
		minFlag = true
		inp = inp * -1
	}
	var fact float64
	fact = math.Gamma(inp+1)
	if minFlag{
		fact = fact * -1
	}
	return fact
}

func (u UnaryExpr) Eval() float64 {
	switch u.op {
	case "!":
		return factorial(u.Operand.Eval())
	case "√":
		return math.Sqrt(u.Operand.Eval())
	case "#": // # is unary minus
		return -1 * u.Operand.Eval()
	}
	return 0
}
