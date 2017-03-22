package generator

import (
	"disownedwheat/gopherscript/parser"
	"fmt"
	"strings"
)

func generator(ast parser.ASTNode) string {

	switch ast.Type {
	case "NEWLINE":
		return ";"
	case "SEMICOLON":
		return ";"
	case "ARGUMENTLISTFUNC":
		var funcBody string
		args := []string{"("}
		var argNumber int
		for i, node := range ast.Body {
			if node.Type == "FUNCTIONBODY" {
				argNumber = i
				break
			}
			args = append(args, node.Value)
		}

		// fmt.Println(ast.Body)

		for _, node := range ast.Body[argNumber:] {
			fmt.Println(node.Value)
			funcBody = generator(node)
		}

		return fmt.Sprintf("func %s%s %s", ast.Value, strings.Join(args, ""), funcBody)

	case "ARGUMENTLIST":
		args := []string{"("}
		for _, node := range ast.Body {
			args = append(args, node.Value)
		}
		return strings.Join(args, "")

	case "FUNCTIONBODY":
		fmt.Println(ast)
		var nodes []string
		for _, node := range ast.Body {
			nodes = append(nodes, generator(node))
		}
		return strings.Join(nodes, "")

	case "IDENT":
		return ast.Value
	case "RPAREN":
		return ")"
	case "LPAREN":
		return "("
	case "LBRACE":
		return "{"
	case "RBRACE":
		return "}"

	case "STRING":
		return fmt.Sprintf("\"%s\"", ast.Value)
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
