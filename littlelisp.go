package main

import "fmt"

type lispValue interface {
  eval() string
  evalLisp() string
}

type Null struct {
}

func null() *Null {
  return new(Null)
}

func (p *Null) eval() string {
  return ""
}

func (p *Null) evalLisp() string {
  return p.eval()
}

type Atom struct {
  value string
}

func (p *Atom) eval() string {
  return p.value
}

func (p *Atom) evalLisp() string {
  return p.eval()
}

func atom(value string) *Atom {
  return &Atom{value}
}

type Pair struct {
  car lispValue
  cdr lispValue
}

func (p *Pair) eval() string {
  if _, ok := p.cdr.(*Null); ok {
    return p.car.eval()
  } else {
    return p.car.evalLisp() + " " + p.cdr.eval()
  }
}

func (p *Pair) evalLisp() string {
  return "(" + p.eval() + ")"
}

func cons(car lispValue, cdr lispValue) lispValue {
  return &Pair{car, cdr}
}

func main() {
  fmt.Println(atom("a").eval())
  fmt.Println(atom("1").eval())

  //'(1 2 3 4)
  list := cons(atom("1"), cons(atom("2"), cons(atom("3"), cons(atom("4"), null()))))
  fmt.Println(list.evalLisp())

  //'(1 (2 3) 4)
  list = cons(atom("1"), cons(cons(atom("2"), cons(atom("3"), null())), cons(atom("4"), null())))
  fmt.Println(list.evalLisp())
}
