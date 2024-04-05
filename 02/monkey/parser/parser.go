package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser has a lexer, and the current and next tokens
type Parser struct {
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token
}

// Can create a parser from a lexer.
func New(l *lexer.Lexer) *Parser {
	// Creates a new instance of Parser and assigns p to it.
	// The & is used to take the address of the newly created Parser.
	// It is used to initialize the Parser.
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

// Errors returns the errors stored on a parser.
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError adds an error whenbt the type of
// the peek token doesn't match the expectation.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expecte next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// nextToken advances the current and next tokens
// of a Parser
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Recursive-descent parsing.
// It starts creating the root node. Then,
// builds the child nodes - statements -
// by calling other functions that know which
// AST node to construct based on the current
// token. These functions can be recursive.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseStatement Checks the current token of a parser to
// classify the type of statement. Then, calls the
// appropriate parsing method for the type of statement.
// It returns a node with the type of the statement.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement Reads a let statement of the form
// Let identifier = expression; and returns the pointer to
// an ast node for the statement.
func (p *Parser) parseLetStatement() *ast.LetStatement {

	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: add the parsing of the expression.

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs checks if the current token of a parser is of type t.
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs checks if the peek token of a parser is of type t.
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek moves the parser's token only if the
// peek token type matches the expectations.
// It's primary purpose it to enforce the correctness
// on the order of tokens.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
