package littlelisp

import "testing"

func TestAtom(t *testing.T) {
  env := NewEnv()

  v, _ := NewString("a").Eval(env)
  if v.String() != "a" {
    t.Errorf("it should be 'a' but was %s", v)
  }

  v, _ = NewNumber(1).Eval(env)
  if v.String() != "1" {
    t.Errorf("it should be '1' but was %s", v)
  }

  env.Define("b", NewNumber(2))
  v, _ = NewSymbol("b").Eval(env)
  if v.String() != "2" {
    t.Errorf("it should be '2' but was %s", v)
  }

  procedure := NewProcedure(func(params *Pair) *Atom {
    return NewNumber(2)
  })
  env.Define("sum", procedure)
  v, _ = NewSymbol("sum").Eval(env)
  if v.String() != "<procedure>" {
    t.Errorf("it should be 'procedure' but was %s", v)
  }
}
