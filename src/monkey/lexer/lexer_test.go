package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	tokens := New(input)

	for i, testCase := range tests {
		token := tokens.NextToken()

		if token.Type != testCase.expectedType {
			t.Fatalf("tests[%d] - wrong token type, expected %q was %q", i, testCase.expectedType, token.Type)
		}

		if token.Literal != testCase.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal, expected %q was %q", i, testCase.expectedLiteral, token.Literal)
		}
	}
}
