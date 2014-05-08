package littlelisp

import (
  . "github.com/littlelisp/nodes"
  "strconv"
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
    return parseItens(lexer)
  }
  return nil
}

func parseItens(lexer *Lexer) *Pair {
  token := lexer.NextToken()

  switch token.typ {
  case tokenRight:
    return EmptyPair()
  case tokenNumber:
    number, _ := strconv.Atoi(token.val)
    return NewPair(NewNumber(number), nil)
  }

  return nil
}
