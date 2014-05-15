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
  {"listWithOneElement", "'(1)", NewPair(NewSymbol("quote"), NewPair(NewPair(NewNumber(1),nil), nil))},
  {"listWithTwoElements", "'(1 \"a\")", NewPair(NewSymbol("quote"),
                                          NewPair(NewPair(NewNumber(1),NewPair(NewString("\"a\""),nil)),nil))},
  {"listWithThreeElements", "'(1 \"a\" s)",
        NewPair(NewSymbol("quote"), NewPair(NewPair(NewNumber(1),
              NewPair(NewString("\"a\""),NewPair(NewSymbol("s"),nil))),nil))},
  {"functionCall1", "(+ 1 2)", NewPair(NewSymbol("+"),
                                          NewPair(NewNumber(1),NewPair(NewNumber(2),nil)))},
  {"functionCall2", "(car (1 2))", NewPair(NewSymbol("car"),
      NewPair(
        NewPair(NewNumber(1),NewPair(NewNumber(2), nil)),
        nil,
        ))},

  {"lambda", "(lambda (a b) (+ a 2))", NewPair( NewSymbol("lambda"),
      NewPair(NewPair(NewSymbol("a"), NewPair(NewSymbol("b"), nil)),
              NewPair(
                NewPair(NewSymbol("+"), NewPair(NewSymbol("a"), NewPair(NewNumber(2), nil))),nil))),
  },
  {"nested", "'((1 2) (3 4) (5 6))", NewPair(
    NewSymbol("quote"),
    NewPair(
      NewPair(NewNumber(1),NewPair(NewNumber(2),nil)),
      NewPair(
        NewPair(NewNumber(3),NewPair(NewNumber(4),nil)),
        NewPair(
          NewPair(NewNumber(5),NewPair(NewNumber(6),nil)), nil),
      ),
    ),
  )},
}

func TestParse(t *testing.T) {
 for _, test := range parseTests {
   got := Parse(test.input)
   if !equalPair(got, test.expected) {
     t.Errorf("%s should be equal, expected %s but was %s", test.name, test.expected, got.String())
   }
 }
}

func equalPair(got, expected Value) bool {
  if got.String() != expected.String() {
    return false
  }
  return true
}
