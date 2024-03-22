package ast

import "monkey/token"

// Statements and Expressions nodes

// The signature of a Node is a token string
type Node interface {
	TokenLiteral() string
}

// The signatures of statments and expressions
// contain dummy methods that will be used to
// throw errors when we use a statements
// where an expression should be used and
// viceversa

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// A let statement has an identifier,
// the expression, and the node of the
// AST associated to it.
type LetStatement struct {
	Name  *Identifier // pointer to the Identifier node
	Value Expression
	Token token.Token // the token.LET token
}


// Go method resolution is based on the receiver type.
// You can define the same method on different types.

// Returns a node of AST??
func (ls *LetStatement) statementNode() {}
// Returns the token literal of a LetStatement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }


// The identifier node contains an IDENT token
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // Identifiers will produce values in other parts of the code
}

// A program is an array of statements
// The program node will be the root of every AST
type Program struct {
	Statements []Statement
}

// Return ??
func (i *Identifier) statementNode() {}
// Returns the token literal of an Identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// TokenLitral stores statements and Expressions
// in a program's Statements and Expressions
// properties.
// Moves through the Statements form the AST
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}


