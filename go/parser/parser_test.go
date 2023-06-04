package parser

import (
  "testing"
  "monkey_compiler/ast"
  "monkey_compiler/lexer"
)

func TestStatements(t *testing.T) {
  input :=
`
let x = 5;
let y = 10;
let foobar = 838383;

`
  
  lex := lexer.New(input)
  prog := New(lex)

  program := prog.ParseProgram()
  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }
  if len(program.Statements) != 3 {
    t.Fatalf("program.Statements does not contain 3 statements. got: %d", len(program.Statements))
  }

  tests := []struct {
    expectedIdentifier string
  }{
    { "x" },
    { "y" },
    { "foobar" },
  }

  for i, tt := range tests {
    stmt := program.Statements[i]
    if !TestLetStatement(t, stmt, tt.expectedIdentifier) {
      return
    }
  }
}

func TestLetStatement(t *testing.T, s ast.Statement, name string) bool {
  if s.TokenLiteral() != "let" {
    t.Errorf("s.TokenLiteral not 'let'. got: %q", s.TokenLiteral())
    return false
  }

  letStmt, ok := s.(*ast.LetStatement)
  if !ok {
    t.Errorf("s not *ast.LetStatement. got: %t", s)
    return false
  }

  if letStmt.Name.TokenLiteral() != name {
    t.Errorf("letStmt.Name.TokenLiteral() not %s. got: %s", name, letStmt.Name.TokenLiteral())
    return false
  }

  return true
}
