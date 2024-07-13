package lexer

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/token"
)

func Scan(current rune) {
	switch current {
	case token.LEFT_PAREN:
		fmt.Println("LEFT_PAREN ( null")
	case token.RIGHT_PAREN:
		fmt.Println("RIGHT_PAREN ) null")
	case token.LEFT_BRACE:
		fmt.Println("LEFT_BRACE { null")
	case token.RIGHT_BRACE:
		fmt.Println("RIGHT_BRACE } null")
	}
}
