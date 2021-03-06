package main

import (
	"disownedwheat/gopherscript/lexer"
	"disownedwheat/gopherscript/parser"
	"disownedwheat/gopherscript/token"
	"disownedwheat/gopherscript/transformer"
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

	input = `
	let z = (x, y) -> {
		let x = 5
		print(x)
	}
	let testFunc = () -> print("This is working!")
`

	lex := []token.Token{}
	l := lexer.New(input)

	for i := 0; i != len(input); i++ {
		tok := l.NextToken()
		if tok.Type != token.EOF {
			lex = append(lex, tok)
		}
	}

	ast := parser.Parse(lex)
	ast2 := transformer.Transformer(ast.Body, &ast)
	fmt.Println(ast2)
	// fmt.Println(parser.Parse(lex))
	// fmt.Println(generator.Generate(parser.Parse(lex)))
}
