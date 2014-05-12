package littlelisp

import (
  . "github.com/littlelisp/nodes"
  "strconv"
)

type Parser struct {
  lexer *Lexer
  tokens [1]token
  tokenIndex int
}

func (p *Parser) nextToken() token {
  if p.tokenIndex > 0 {
    p.tokenIndex--
    return p.tokens[p.tokenIndex]
  } else {
    t := p.lexer.NextToken()
    p.tokens[p.tokenIndex] = t
    p.tokenIndex = 0
    return t
  }
}

func (p *Parser) backup() {
  p.tokenIndex++
}


func Parse(input string) *Pair {
  lexer := Lex(input)
  parser := &Parser{lexer: lexer, tokenIndex: 0}
  return parsePair(parser)
}

func parsePair(parser *Parser) *Pair {
  token := parser.nextToken()

  switch token.typ {
  case tokenQuote:
    return NewPair(NewSymbol("quote"), parsePair(parser))
  case tokenLeft:
    token := parser.nextToken()
    switch token.typ {
    case tokenRight:
      return EmptyPair()
    case tokenSymbol:
      return NewPair(NewSymbol(token.val), parsePair(parser))
    default:
      parser.backup()
      return parsePair(parser)
    }
  case tokenRight:
    return nil
  case tokenNumber:
    number, _ := strconv.Atoi(token.val)
    return NewPair(NewNumber(number), parsePair(parser))
  case tokenString:
    return NewPair(NewString(token.val), parsePair(parser))
  case tokenSymbol:
    return NewPair(NewSymbol(token.val), parsePair(parser))
  }

  return nil
}
