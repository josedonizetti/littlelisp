package littlelisp

import(
  "testing"
  "fmt"
)

var (
    tEOF = token{tokenEOF,""}
    tQuote = token{tokenQuote,"'"}
    tLeft = token{tokenLeft,"("}
    tRight = token{tokenRight,")"}
)

func equal(got, exp []token) bool {

  if len(got) != len(exp) {
    return false
  }

  for index, _ := range got {
    if got[index].typ != exp[index].typ {
      fmt.Println(got[index].typ)
      fmt.Println(exp[index].typ)
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

type LexTest struct {
  name string
  input string
  tokens []token
}

var tests = []LexTest{
  {"emptyList", "'()", []token{
    tQuote,
    tLeft,
    tRight,
    tEOF,
  }},
  {"listWithBasicNumber", "'(1)", []token{
    tQuote,
    tLeft,
    {tokenNumber,"1"},
    tRight,
    tEOF,
  }},
  {"listWithComplexNumber", "'(1011)", []token{
    tQuote,
    tLeft,
    {tokenNumber,"1011"},
    tRight,
    tEOF,
  }},
  {"listWithNumberAndString", "'(1011 \"string\")", []token{
    tQuote,
    tLeft,
    {tokenNumber,"1011"},
    {tokenString,"\"string\""},
    tRight,
    tEOF,
  }},
  {"nestedList", "'(1011 (\"string\" 1) \"another\")", []token{
    tQuote,
    tLeft,
    {tokenNumber,"1011"},
    tLeft,
    {tokenString,"\"string\""},
    {tokenNumber,"1"},
    tRight,
    {tokenString,"\"another\""},
    tRight,
    tEOF,
  }},
}

func TestLex(t *testing.T) {
  for _,test := range tests {
    lexer := lex(test.input)
    tokens := collect(lexer)
    if !equal(tokens, test.tokens) {
      t.Errorf("Spec %s should be equal, but got %v and was expected %v", test.name, tokens, test.tokens)
    }
  }
}
