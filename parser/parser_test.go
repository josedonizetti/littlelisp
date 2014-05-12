package parser

import (
  "testing"
  . "github.com/littlelisp/context"
)

type ParseTest struct {
  name string
  input string
  expected Value
}

var parseTests = []ParseTest{
  {"number", "1", NewNumber(1)},
  {"symbol", "symbol", NewSymbol("symbol")},
  {"string", "\"string\"", NewString("\"string\"")},
  {"emptyList", "'()", NewPair(NewSymbol("quote"),EmptyPair())},
  {"listWithOneElement", "'(1)", NewPair(NewSymbol("quote"),NewPair(NewNumber(1),nil))},
  {"listWithTwoElements", "'(1 \"a\")", NewPair(NewSymbol("quote"),
                                          NewPair(NewNumber(1),NewPair(NewString("\"a\""),nil)))},
  {"listWithThreeElements", "'(1 \"a\" s)",
        NewPair(NewSymbol("quote"), NewPair(NewNumber(1),
              NewPair(NewString("\"a\""),NewPair(NewSymbol("s"),nil))))},
  {"functionCall", "(car (1 2))", NewPair(NewSymbol("car"),
                                          NewPair(NewNumber(1),NewPair(NewNumber(2),nil)))},
}

func TestParse(t *testing.T) {
 for _, test := range parseTests {
   got := Parse(test.input)
   if !equalPair(got, test.expected) {
     t.Errorf("%s should be equal, expected %s but was %s", test.name, test.expected.String(), got.String())
   }
 }
}

func equalPair(got, expected Value) bool {
  if got.String() != expected.String() {
    return false
  }
  return true
}