package lexer

import "monkey/token"

/*
  Lexer type contains information about the input
  that is being converter into tokens.
*/
type Lexer struct {
	input        string //
	position     int    // current position on the input
	readPosition int    // current reading position. Always points to the next character from position.
	ch           byte   // current character under examination. Only ASCII characters are supported.
}

// Returns a new instance of the struct type Lexer
// for an input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // sets ch to 0 and readPosition to 1.
	return l
}

// readChar gives the next character and updates both positions.
// ch is set to 0 (ASCII for null) if you:
//	- haven't read anything yet. (Used in New).
//	- have reached the end of the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = '0'
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// peekChar returns the next character from the one under examination.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return '0'
	} else {
		return l.input[l.readPosition]
	}
}

// NextToken returns the token defined by the character
// under examination of a Lexer instance and moves to the Lexer
// to the next character.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace() // white space is only used to separate tokens.

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '0':
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // return here because readIdentifier moves the Lexer to the end of the identifier.
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok // return here because readIdentifier moves the Lexer to the end of the identifier.
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readIdentifier returns the sequence of letters ([a-zA-z_]) that starts from
// the current character under examination of a Lexer.
// It moves the Lexer to the end of the sequence.
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// isLetter checks if an argument is a letter (i.e. in [a-zA-z_]).
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// readNumber returns the sequence of digits that starts from
// the current character under examination of a Lexer.
// It moves the Lexer to the end of the sequence.
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// isDigit checks if an argument is [0-9].
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// newToken returns a token instance given the type and
// source code ASCII character.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// skipWhitespace moves the Lexer when the token is whitespace.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

