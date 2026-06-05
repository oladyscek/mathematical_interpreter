package main

import (
	"strings"
	"unicode"
)

func commaLexer(st_input string) string {
	sl := make([]rune, 0, len(st_input))
	for _, val := range st_input {
		if val == ',' { // comma to dot
			sl = append(sl, '.')
			continue
		}
		if val == ' ' { //ignore spaces
			continue
		} else { //just normal num
			sl = append(sl, val)
		}
	}
	str := string(sl)
	return str
}

func validLexer(input string) bool {
	var validOps = map[rune]struct{}{ // valid letters except numbers
		'+': {}, '-': {}, '*': {}, '/': {},
		'(': {}, ')': {}, '^': {}, ' ': {}, '#': {},
		'!': {}, '√': {}, 'π': {}, '.': {}, ',': {},
	}
	for _, val := range input {
		if _, ok := validOps[val]; ok || unicode.IsDigit(val) { // if number or another valid letter
			continue // check next symbol
		} else {
			return false // if not valid
		}
	}
	return true
}

func unaryLexer(input []string) []string {
	sls := make([]string, 0, len(input)) 
	i := 0
	var unary = map[string]struct{}{ // set of unary triggers (before unary minus)
		"+": {}, "-": {}, "*": {}, "/": {},
		"(": {}, "^": {},
	}
	for i < len(input) {
		if i == 0 && input[i] == "-" { // if first is minus => unary
			sls = append(sls, "#")
			i++
			continue // go to next token
		}
		if i > 0 && input[i] == "-" { // index could be still 0 and [input[i-1]] will end with an error
			if _, ok := unary[input[i-1]]; ok {
				sls = append(sls, "#") // # is unary minus
				i++
				continue // go to next token
			}
			if _, ok := unary[input[i-1]]; ok { // if minus and token before in unary triggers
				var sb strings.Builder
				sb.WriteString("-") // building a new token
				sb.WriteString(input[i+1])
				sls = append(sls, sb.String()) // add new token
				i += 2
				continue // go to next token
			}
		} 
		sls = append(sls, input[i]) // if num just add and go to next token
		i++
	}

	ind := 0
	res := make([]string, 0, len(sls)) // new slice without unary pluses
	for ind < len(sls) {
		if ind == 0 && sls[ind] == "+" {
			ind++ // if unary plus just skip it, unary pluses are trash
			continue // go to next token
		}
		if ind > 0 && sls[ind] == "+" { // index could be still 0 and [sls[ind-1]] will end with an error
			if _, ok := unary[sls[ind-1]]; ok { // if unary somwhere in the middle in the tokens
				ind++ // skip it again
				continue
			}
		}
		res = append(res, sls[ind])
		ind++
	}
	return res
}


func lexer(input string) []string {
	tokens := make([]string, 0, len(input))
	index := 0
	if validLexer(input) == false {
		return []string{"error"} // sorry for that
	}
	st_given := commaLexer(input)
	runes := []rune(st_given)
	for index < len(runes) { //while index is not out of range
		if unicode.IsDigit(runes[index]) || runes[index] == '.' { // if num or dot => num = true
			start := index
			for index < len(runes) && (unicode.IsDigit(runes[index]) || runes[index] == '.') { //find index of num
				index++
			}
			tokens = append(tokens, string(runes[start:index])) // add the num to slice
		} else {
			tokens = append(tokens, string(runes[index])) // if not num just add
			index++
		}
	}
	return unaryLexer(tokens)
}