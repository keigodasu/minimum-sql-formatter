package minimum_sql_formatter

import "strings"

type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

func NewLexer(input string) *Lexer  {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar()  {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case ',':
		tok = newToken(COMMA, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier()  string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return strings.ToUpper(l.input[position:l.position])
}

func (l *Lexer) readNumber() string  {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool{
	return 'a'	<= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || '0' <= ch && ch <= '9'
}

func isDigit(ch byte) bool  {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace()  {
	for l.ch == ' '	 || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}