package littlelisp

import "testing"

func TestPair(t *testing.T) {
  list := EmptyPair()

  if list.String() != "()" {
    t.Errorf("it should be equal, expecte () got %s", list.String())
  }

  list = NewPair(NewString("1"),NewPair(NewNumber(1),nil))
  if list.String() != "(\"1\" 1)" {
    t.Errorf("it should be equal, expecte (\"1\" 1) got %s", list.String())
  }
}
