package littlelisp

import (
  "testing"
  . "github.com/littlelisp/nodes"
)

type ParseTest struct {
  name string
  input string
  expected *Pair
}

var parseTests = []ParseTest{
  {"emptyList", "'()", NewPair(NewSymbol("quote"),EmptyPair())},
  {"listWithOneNumber", "'(1)", NewPair(NewSymbol("quote"),NewPair(NewNumber(1),nil))},
}

func TestParse(t *testing.T) {
 for _, test := range parseTests {
   got := Parse(test.input)
   if !equalPair(got, test.expected) {
     t.Errorf("%s should be equal, expected %s but was %s", test.name, test.expected.String(), got.String())
   }
 }
}

func equalPair(got, expected *Pair) bool {
  if got.String() != expected.String() {
    return false
  }
  return true
}
