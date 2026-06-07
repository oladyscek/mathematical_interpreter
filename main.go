package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		given := bufio.NewScanner(os.Stdin)
		fmt.Println("enter expression (please)")
		given.Scan()
		input := given.Text()
		tokens := lexer(input)
		ast := parse(tokens)
		res := ast.Eval()
		fmt.Println(format(res))
	}
}
