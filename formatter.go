package main

import (
	"fmt"
	"math"
)

func reverseString(input string) string {
	sl := make([]rune, 0, len(input))
	i := len(input) - 1 // index of last num
	for i != -1 {       // we also need token with index 0, when i-- on the last iteration we get i = -1
		sl = append(sl, rune(input[i]))
		i--
	}
	return string(sl)
}

func commaFormatter(input string) string {
	list := make([]rune, 0, len(input))
	for _, tok := range input {
		if tok == '.' { // if dot add the comma
			list = append(list, ',')
		} else { // if not dot => number. just add it
			list = append(list, tok)
		}
	}
	return string(list)
}

func format(input float64) string {
	if input == math.Trunc(input) { // if we get 5 or 5.000000 for example we do not need to format it
		return fmt.Sprintf("%.0f", input) // "%.0f" do from 5.00000 to 5
	}
	str := reverseString(fmt.Sprintf("%.14f", input)) // reverse our string to cut some 0 which we do not need
	sls := make([]rune, 0, len(str))

	mem := -1 // index of first important token
	for ind, tok := range str {
		if tok == '0' { // zero which we do not need
			continue
		} else { // if not zero we need it
			mem = ind // index of first important token
			break
		}
	}

// remember that str is reversed, it looks like 000007654.321

	for mem != len(str) { // add tokens which we need in the new slice
		sls = append(sls, rune(str[mem]))
		mem++
	}

	ans := reverseString(string(sls)) // make string and reverse back
	return commaFormatter(ans) // replace "." to ","
}
