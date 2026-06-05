package main

import "strconv"

func pres(op string) int {
	switch op {
	case "+", "-": // have the lowest pres
		return 1
	case "*", "/": // have the higher pres
		return 2
	case "#":
		return 3
	case "^", "√", "!": // have the highest pres
		return 4
	}
	return 9999
}

func isOperator(inp string) bool {
	var operators = map[string]struct{}{
		"+": {}, "-": {}, "*": {}, "/": {},
		"^": {}, "!": {}, "√": {}, "#": {},
	}
	_, ok := operators[inp] // ok returns true if there is such key in operators
	return ok
}

func root(input []string) int {
	lowestpres := 9999
	root_pres := -1
	depth := 0 // operator in () is not root

	for i, tok := range input {
		if tok == "(" { // if we are inside of brackets go untill we are out of them
			depth++
			continue
		}
		if tok == ")" {
			depth--
			continue
		}
		if depth == 0 { // if we are not in brackets 
			if isOperator(tok) {
				prec := pres(tok)       // pres of current operator
				if prec <= lowestpres { // if lower
					lowestpres = prec
					root_pres = i
				}
			}
		}
	}
	return root_pres
}

func IsUselessSk(input []string) bool {
	if len(input) == 0 {
		return false
	}
	depth := 0
	if input[0] == "(" && input[len(input)-1] == ")" { // it could be case (2+2), we need to check it
		for i, val := range input {
			if val == "(" {
				depth++ // if the opening bracket we are inside brackets
				continue
			}
			if val == ")" {
				depth--
				if depth == 0 { // if we are not in brackets
					if i == len(input)-1 { // if we are not in brackets and there is end
						return true // if that case we throw out brackets. (2+2) => 2+2
					}
					return false // if we get out of brackets and it's not end, it's not our case
				}
			}
		}
	}
	return false // if there is no opening bracket
}

func parse(input []string) Node {
	if len(input) == 1 { // if just 1 num return it
		val, _ := strconv.ParseFloat(input[0], 64)
		return Number{
			Value: val,
		}
	}

	if IsUselessSk(input) { // could be (2+2) and that will break everything
		return parse(input[1 : len(input)-1]) // if that case just throw out useless brackets
	}

	root_ind := root(input)
	if root_ind == -1 { // if no root found 
		panic("invalid expression")
	}
	switch input[root_ind] {
	case "+", "-", "/", "*", "^":
		return BinaryExpr{
			op:    input[root_ind],
			Left:  parse(input[:root_ind]),
			Right: parse(input[root_ind+1:]),
		}
	case "!":
		return UnaryExpr{
			op:      "!",
			Operand: parse(input[:root_ind]), // postfix operand 
		}
	case "√", "#":
		return UnaryExpr{
			op:      input[root_ind],
			Operand: parse(input[root_ind+1:]), // prefix operand
		}
	}
	panic("unknown operator")
}
