package parser

import(
  "testing"
)

var (
    tEOF = token{tokenEOF,""}
    tQuote = token{tokenQuote,"'"}
    tLeft = token{tokenLeft,"("}
    tRight = token{tokenRight,")"}
)

func equalTokens(got, exp []token) bool {

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
    token := lexer.NextToken()
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

var lexTests = []LexTest{
  {"number", "1", []token{
    {tokenNumber,"1"},
    tEOF,
  }},
  {"symbol", "a", []token{
    {tokenSymbol,"a"},
    tEOF,
  }},
  {"string", "\"a\"", []token{
    {tokenString,"\"a\""},
    tEOF,
  }},
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
  {"functionCall1", "(+ 1 2)", []token{
    tLeft,
    {tokenSymbol,"+"},
    {tokenNumber,"1"},
    {tokenNumber,"2"},
    tRight,
    tEOF,
  }},
  {"functionCall2", "(car (1 2 \"string\"))", []token{
    tLeft,
    {tokenSymbol,"car"},
    tLeft,
    {tokenNumber,"1"},
    {tokenNumber,"2"},
    {tokenString,"\"string\""},
    tRight,
    tRight,
    tEOF,
  }},
  {"lambdaCall", "(lambda (a b) (+ a 2))", []token{
    tLeft,
    {tokenSymbol,"lambda"},
    tLeft,
    {tokenSymbol,"a"},
    {tokenSymbol,"b"},
    tRight,
    tLeft,
    {tokenSymbol,"+"},
    {tokenSymbol,"a"},
    {tokenNumber,"2"},
    tRight,
    tRight,
    tEOF,
  }},
}

func TestLex(t *testing.T) {
  for _,test := range lexTests {
    lexer := Lex(test.input)
    tokens := collect(lexer)
    if !equalTokens(tokens, test.tokens) {
      t.Errorf("Spec %s should be equal, but got %v and was expected %v", test.name, tokens, test.tokens)
    }
  }
}
