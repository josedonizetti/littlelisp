package littlelisp

import (
  "fmt"
  "testing"
)

func TestLittleLisp(t *testing.T) {
  fmt.Println(atom("a").eval())
  fmt.Println(atom("1").eval())

  //'()
  fmt.Println(emptyList().evalList())

  //'(1 2 3 4)
  list := cons(atom("1"), cons(atom("2"), cons(atom("3"), cons(atom("4"), null()))))
  fmt.Println(list.evalList())

  //'(1 (2 3) 4)
  list = cons(atom("1"), cons(cons(atom("2"), cons(atom("3"), null())), cons(atom("4"), null())))
  fmt.Println(list.evalList())
}
