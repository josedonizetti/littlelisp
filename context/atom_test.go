package context

import "testing"

func TestAtom(t *testing.T) {
  env := NewEnv(nil)
  forms := GetForms()

  v, _ := NewString("\"a\"").Eval(env, forms)
  if v.String() != "\"a\"" {
    t.Errorf("it should be \"a\" but was %s", v)
  }

  v, _ = NewNumber(1).Eval(env, forms)
  if v.String() != "1" {
    t.Errorf("it should be '1' but was %s", v)
  }

  v = NewSymbol("a")
  if v.String() != "a" {
    t.Errorf("it should be a but was %s", v)
  }

  env.Define("b", NewNumber(2))
  v, _ = NewSymbol("b").Eval(env, forms)
  if v.String() != "2" {
    t.Errorf("it should be '2' but was %s", v)
  }
}
