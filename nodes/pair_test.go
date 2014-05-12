package littlelisp

import "testing"

func TestPair(t *testing.T) {
  list := EmptyPair()

  if list.String() != "()" {
    t.Errorf("it should be equal, expected () got %s", list.String())
  }

  list = NewPair(NewString("\"1\""),NewPair(NewNumber(1),nil))
  if list.String() != "(\"1\" 1)" {
    t.Errorf("it should be equal, expected (\"1\" 1) got %s", list.String())
  }

  list = NewPair(NewSymbol("quote"),EmptyPair())
  if list.String() != "(quote ())" {
    t.Errorf("it should be equal, expected (quote ()) got %s", list.String())
  }
}
