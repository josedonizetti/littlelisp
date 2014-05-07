package littlelisp

import "testing"

func TestPair(t *testing.T) {
  list := NewPair(NewNumber(1),NewPair(NewNumber(2), nil))
  if list.String() != "(1 2)" {
    t.Errorf("should be (1 2) but was %s", list.String())
  }

  env := NewEnv()
  env.Define("+", NewProcedure(Add))
  list = NewPair(NewSymbol("+"), NewPair( NewNumber(1), NewPair(NewNumber(2), nil)))
  atom := list.Eval(env)
  if atom.String() != "3" {
    t.Errorf("should be 3 but was %s", atom.String())
  }
}
