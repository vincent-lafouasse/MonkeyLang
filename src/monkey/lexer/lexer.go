package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var output_token token.Token

	switch l.ch {
	case '=':
		output_token = newToken(token.ASSIGN, l.ch)
	case ';':
		output_token = newToken(token.SEMICOLON, l.ch)
	case '(':
		output_token = newToken(token.LPAREN, l.ch)
	case ')':
		output_token = newToken(token.RPAREN, l.ch)
	case '{':
		output_token = newToken(token.LBRACE, l.ch)
	case '}':
		output_token = newToken(token.RBRACE, l.ch)
	case ',':
		output_token = newToken(token.COMMA, l.ch)
	case '+':
		output_token = newToken(token.PLUS, l.ch)
	case 0:
		output_token.Type = token.EOF
		output_token.Literal = ""
	default:
		if isLetter(l.ch) {
			output_token.Literal = l.readIdentifier()
			output_token.Type = LookupIdentifier(output_token.Literal)
			return output_token
		} else {
			output_token = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()

	return output_token
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	start_position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[start_position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}
