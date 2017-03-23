package transformer

import (
	"disownedwheat/gopherscript/parser"
	"disownedwheat/gopherscript/token"
)

func Transformer(body []parser.ASTNode, parent *parser.ASTNode) []parser.ASTNode {

	var newBody []parser.ASTNode

	i := 0
	for i != len(body) {
		node := body[i]

		// LET STATEMENTS
		switch node.Type {
		case token.LET:
			if body[i+3].Type == token.FUNCTION || body[i+3].Type == "ARGUMENTLISTFUNC" {
				// If this is a let expression, set the value to the identifier
				var Type token.TokenType
				if parent.Type == "Program" {
					Type = "TOPLEVELFUNCDEFINITION"
				} else {
					Type = "FUNCDEFINITION"
				}
				newNode := parser.ASTNode{
					Type:  Type,
					Value: body[i+1].Value,
					Body:  Transformer(body[i+3].Body, &body[i+3]),
				}
				i = i + 2
				newBody = append(newBody, newNode)
			} else {
				newBody = append(newBody, parser.ASTNode{
					Type:  "LETEXPRESSION",
					Value: body[i+1].Value,
					Body:  []parser.ASTNode{},
				})
				i++
			}
			break
		case token.IDENT:
			if parent.Type == "ARGUMENTLISTFUNC" {
				thisNode := parser.ASTNode{
					Type:  "LPAREN",
					Value: "(",
					Body:  []parser.ASTNode{},
				}
				newBody = append(newBody, thisNode)

				i++
				continue
			}
			if body[i-1].Type == token.LET {
				i++
				continue
			}

			i++
			break
		case "ARGUMENTLISTFUNC":
			newNode := parser.ASTNode{
				Type:  "FUNCTIONDEFINITION",
				Value: "",
				Body:  Transformer(node.Body, &node),
			}

			newBody = append(newBody, newNode)
			i++
			break
		default:
			newBody = append(newBody, node)
			i++
			break
		}
	}

	return newBody
}
