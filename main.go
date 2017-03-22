package main

import (
	"disownedwheat/gopherscript/lexer"
	"disownedwheat/gopherscript/parser"
	"disownedwheat/gopherscript/token"
	"fmt"
)

func main() {
	input := `
let x = 1
let y = [1, 2, 3, 4]
let z = (x, y) -> {
	print(x);
}
let testFunc = () -> print("This is working!")
let value = do -> return x
	`

	lex := []token.Token{}
	l := lexer.New(input)

	for i := 0; i != len(input); i++ {
		tok := l.NextToken()
		if tok.Type != token.EOF {
			lex = append(lex, tok)
		}
	}

	fmt.Println(lex)
	fmt.Println(parser.Parse(lex))

}
