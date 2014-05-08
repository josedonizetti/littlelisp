package littlelisp

import (
  . "github.com/littlelisp/nodes"
/*  "fmt"*/
)

func Parse(input string) *Pair {
  lexer := Lex(input)
  token := lexer.NextToken()

  switch token.typ {
  case tokenQuote:
    return NewPair(NewSymbol("quote"), parsePair(lexer))
  }

  return new(Pair)
}

func parsePair(lexer *Lexer) *Pair {
  token := lexer.NextToken()
  switch token.typ {
  case tokenLeft:
    token = lexer.NextToken()
    if token.typ == tokenRight {
      return EmptyPair()
    }
  }
  return nil
}
