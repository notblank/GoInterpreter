package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser has a lexer, and the current and next tokens
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// Can create a parser from a lexer.
func New(l *lexer.Lexer) *Parser {
	// Creates a new instance of Parser and assigns p to it.
	// The & is used to take the address of the newly created Parser.
	// It is used to initialize the Parser.
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

// nextToken advances the current and next tokens
// of a Parser
func (p *Parser) nextToken(){
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
		return nill
	}
}


// parseLetStatement Reads a let statement of the form
// Let identifier = expression; and returns an ast
// node for the statement.
func (p *Parser) parseLetStatement() ast.LetStatement {

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


