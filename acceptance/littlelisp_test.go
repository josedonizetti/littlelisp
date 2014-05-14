package littlelisp

import (
  "testing"
  . "github.com/littlelisp/context"
)

func TestLittleLisp(testing *testing.T) {
  env := NewEnv(nil)
  forms := GetForms()

  //"a"
  NewString("a").Eval(env, forms)

  // 1
  NewNumber(1).Eval(env, forms)

  //'()
  NewPair(nil,nil).Eval(env, forms)

  //'(1 2 3)
  NewPair(NewNumber(1), NewPair(NewNumber(2), NewPair(NewNumber(3), nil))).Eval(env, forms)

  //(define a 2)
  env.Define("a", NewNumber(2))

  //a
  NewSymbol("a").Eval(env, forms)

  //'(1 a 3)
  NewPair(NewNumber(1), NewPair(NewSymbol("a"), NewPair(NewNumber(3), nil))).Eval(env, forms)

  //(quote a 3)
  NewPair(NewSymbol("quote"), NewPair(NewSymbol("a"), NewPair(NewNumber(3), nil))).Eval(env, forms)

  //(1 a 3) -> error
  NewPair(NewSymbol("1"), NewPair(NewSymbol("a"), NewPair(NewNumber(3), nil))).Eval(env, forms)

  //(+ 2 3)
  NewPair(NewSymbol("+"), NewPair(NewNumber(2), NewPair(NewNumber(3), nil))).Eval(env, forms)

  //(lambda (a b) (+ a b))(1 2)

  //(define sum (lambda (a b) (+ a b)))
  //(sum 1 2)
}
