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


/*
function parseProgram() {
	program = newProgramASTNode()
	advanceTokens()
	for (currentToken() != EOF_TOKEN) {
	statement = null

	if (currentToken() == LET_TOKEN) {
		statement = parseLetStatement()
	} else if (currentToken() == RETURN_TOKEN) {
		statement = parseReturnStatement()
	} else if (currentToken() == IF_TOKEN) {
		statement = parseIfStatement()
	}

	if (statement != null) {
		program.Statements.push(statement)
	}

		advanceTokens()
	} // For loop ending curly brace.

	return program
}
 */

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.ParseStatement()
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
		return p.parseLetStatements()
	default:
		return nil
	}
}



/* Parsing let statements pseudo code.

	function parseLetStatement() {
		advanceTokens()
  		identifier = parseIdentifier()

		advanceTokens()

		if currentToken() != EQUAL_TOKEN {
			parseError("no equal sign!") return null
		}

		advanceTokens()

		value = parseExpression()
		variableStatement = newVariableStatementASTNode()
		variableStatement.identifier = identifier
		variableStatement.value = value
		return variableStatement
}
 */

func (p *Parser) parseLetStatements() *ast.LetStatement {

	// constructs an *ast.LetStatement node with the token it’s currently sitting on (a token.LET token)
	stmt  := &ast.LetStatement{ Token: p.curToken }

	// assert that the next token is an Identifier.
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// construct an *ast.Identifier node.
	stmt.Name = &ast.Identifier{ Token: p.curToken,  Value: p.curToken.Literal }

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

