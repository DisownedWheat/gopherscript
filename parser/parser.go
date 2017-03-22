package parser

import (
	"disownedwheat/gopherscript/token"
	"fmt"
)

type ASTNode struct {
	Type  token.TokenType
	Value string
	name  string
	Body  []ASTNode
}

var currentPos = 0

func walk(input []token.Token) ASTNode {
	tok := input[currentPos]

	if tok.Type == token.ILLEGAL {
		panic(fmt.Sprintf("UNIDENTIFIED TOKEN: %v", tok))
	}

	if tok.Type == token.NEWLINE {
		currentPos++
		return ASTNode{
			Type:  "NEWLINE",
			Value: ";",
			Body:  []ASTNode{},
		}
	}

	if tok.Type == token.SEMICOLON {
		currentPos++
		return ASTNode{
			Type:  "SEMICOLON",
			Value: ";",
			Body:  []ASTNode{},
		}
	}

	if tok.Type == token.INT {
		currentPos++
		return ASTNode{
			Type:  "INT",
			Value: tok.Literal,
			Body:  []ASTNode{},
		}
	}

	if tok.Type == token.STRING {
		currentPos++
		return ASTNode{
			Type:  "STRING",
			Value: tok.Literal,
			Body:  []ASTNode{},
		}
	}

	// Will this be argument list for functions? Let's see
	if tok.Type == token.LPAREN && input[currentPos-1].Type != token.IDENT {
		var name string
		if input[currentPos-1].Type == token.ASSIGN {
			name = input[currentPos-2].Literal
		} else {
			name = ""
		}
		currentPos++
		var body []ASTNode
		for input[currentPos].Type != token.RPAREN {
			body = append(body, walk(input))
			tok = input[currentPos]
		}
		body = append(body, walk(input))

		if input[currentPos].Type == token.FUNCTION {
			body = append(body, walk(input))
		}
		return ASTNode{
			Type:  "ARGUMENTLISTFUNC",
			Value: name,
			Body:  body,
		}
	}
	if tok.Type == token.LPAREN && input[currentPos-1].Type == token.IDENT {
		currentPos++
		var body []ASTNode
		for input[currentPos].Type != token.RPAREN {
			body = append(body, walk(input))
			tok = input[currentPos]
		}
		return ASTNode{
			Type:  "ARGUMENTLIST",
			Value: "",
			Body:  body,
		}
	}

	if tok.Type == token.FUNCTION {
		// Get the previous item (argument list) as well
		currentPos++
		var body = []ASTNode{}
		if input[currentPos].Type != token.LBRACE {
			for input[currentPos].Type != token.NEWLINE {
				body = append(body, walk(input))
			}
		} else {
			for input[currentPos].Type != token.RBRACE {
				body = append(body, walk(input))
			}

		}
		return ASTNode{
			Type:  "FUNCTIONBODY",
			Value: "func",
			Body:  body,
		}
	}

	if tok.Type == token.LET {
		currentPos++
		return ASTNode{
			Type:  "LET",
			Value: "let",
			Body:  []ASTNode{},
		}
	}

	if tok.Type == token.IDENT {
		currentPos++
		return ASTNode{
			Type:  "IDENT",
			Value: tok.Literal,
			Body:  []ASTNode{},
		}
	}

	if tok.Type == token.ASSIGN {
		currentPos++
		return ASTNode{
			Type:  "ASSIGN",
			Value: tok.Literal,
			Body:  []ASTNode{},
		}
	}

	if tok.Type == token.DO {
		currentPos++
		var body []ASTNode
		if input[currentPos].Type != token.LBRACE {
			for input[currentPos].Type != token.NEWLINE {
				body = append(body, walk(input))
			}
		} else {
			for input[currentPos].Type != token.RBRACE {
				body = append(body, walk(input))
			}
		}

		return ASTNode{
			Type:  "DO",
			Value: tok.Literal,
			Body:  body,
		}
	}

	if tok.Type == token.LPAREN {
		currentPos++
		return ASTNode{
			Type:  "LPAREN",
			Value: tok.Literal,
			Body:  []ASTNode{},
		}
	}

	if tok.Type == token.RPAREN {
		currentPos++
		return ASTNode{
			Type:  "RPAREN",
			Value: tok.Literal,
			Body:  []ASTNode{},
		}
	}

	currentPos++
	return ASTNode{
		Type:  "GENERIC",
		Value: tok.Literal,
		Body:  []ASTNode{},
	}
}

func Parse(input []token.Token) ASTNode {

	var program ASTNode
	program.Type = "Program"
	program.Value = ""
	program.Body = []ASTNode{}

	for currentPos != len(input) {
		program.Body = append(program.Body, walk(input))
	}
	return program
}
