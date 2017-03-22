package generator

import (
	"disownedwheat/gopherscript/parser"
	"strings"
)

func generator(ast parser.ASTNode) string {

	switch ast.Type {
	case "NEWLINE":
		return ";"
	case "SEMICOLON":
		return ";"
	default:
		return ""
	}
}

func Generate(ast parser.ASTNode) string {
	var code []string

	for _, node := range ast.Body {
		code = append(code, generator(node))
	}

	return strings.Join(code, "")
}
