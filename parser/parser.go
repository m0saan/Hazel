package parser

import (
	"Interpreter/ast"
	"Interpreter/lexer"
	"Interpreter/token"
)

type Parser struct {
	l *lexer.Lexer // l is an instance of the lexer on which we call NextToken to get the next token in the input

	// both curToken and peekToken point to the current and next token
	curToken token.Token // Token under examination
	peekToken token.Token // Use peekToken to decide what to do next
}

func (p *Parser) nextToken()  {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.ParseProgram()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) ParseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.ParseStatement()
	default:
		return nil
	}
}

