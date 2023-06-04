package parser

import (
  "monkey_compiler/ast"
  "monkey_compiler/lexer"
  "monkey_compiler/token"
)

type Parser struct {
  lex *lexer.Lexer

  curToken token.Token
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
  prog := &Parser{lex: l}

  prog.lex.NextToken()

  prog.lex.NextToken()

  return prog
}

func (p *Parser) ParseProgram() *ast.Program {
  return nil
}
