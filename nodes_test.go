package littlelisp

import (
  "fmt"
  "testing"
)

func TestLittleLisp(t *testing.T) {
  env := DefaultEnv()

  fmt.Println(atom("a").eval(env))
  fmt.Println(atom("1").eval(env))

  //'()
  fmt.Println(emptyList().evalList(env))

  //'(1 2 3 4)
  list := cons(atom("1"), cons(atom("2"), cons(atom("3"), cons(atom("4"), null()))))
  fmt.Println(list.evalList(env))

  //'(1 (2 3) 4)
  list = cons(atom("1"), cons(cons(atom("2"), cons(atom("3"), null())), cons(atom("4"), null())))
  fmt.Println(list.evalList(env))

  //(+ 3 4)
  list = cons(symbol("+"), cons(atom("1"), atom("2")))
  fmt.Println(list.evalList(env))
}
