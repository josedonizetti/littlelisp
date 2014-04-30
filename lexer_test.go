package littlelisp

import(
  "testing"
  "fmt"
)

var (
    tEOF = token{tokenEOF,""}
)

func equal(got, exp []token) bool {

  if len(got) != len(exp) {
    return false
  }

  for index, _ := range got {
    if got[index].typ != exp[index].typ {
      return false
    }

    if got[index].val != exp[index].val {
      return false
    }
  }

  return true
}

func collect(lexer *Lexer) []token {
  tokens := make([]token,0)
  for {
    token := lexer.nextToken()
    tokens = append(tokens,token)

    if token.typ == tokenEOF {
      break
    }
  }
  return tokens
}

func TestLex(t *testing.T) {
  expected := []token{
    {tokenLeft, "("},
    {tokenText, "1"},
    {tokenRight, ")"},
    tEOF,
  }

  lexer := lex("(1)")
  tokens := collect(lexer)

  if !equal(tokens, expected) {
    fmt.Println(tokens)
    fmt.Println(expected)
    t.Errorf("it should be equal, but got %v and was expected %v", tokens, expected)
  }
}
