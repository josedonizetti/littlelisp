package littlelisp

import (
  "strconv"
)

type valueType uint8

const (
  nilValue valueType = iota
  numberValue
  stringValue
  symbolValue
)

type Atom struct {
  typ valueType
  val interface{}
}

var Nil = &Atom{nilValue,nil}

func (a *Atom) Eval(env *Env) (*Atom, error) {
  switch a.typ {
  case symbolValue:
    return env.Lookup(a.val.(string)), nil
  default:
    return a, nil
  }
}

func (a *Atom) String() string {
  switch a.typ {
  case symbolValue:
    return a.val.(string)
  case stringValue:
    return "\"" + a.val.(string) + "\""
  case numberValue:
    return strconv.Itoa(a.val.(int))
  default:
    return ""
  }
}

func (a *Atom) Number() int {
  return a.val.(int)
}

func (a *Atom) IsSymbol() bool {
  return a.typ == symbolValue
}

func NewAtom(val interface{},typ valueType) *Atom {
  return &Atom{typ,val}
}

func NewString(val string) *Atom {
  return &Atom{stringValue,val}
}

func NewNumber(val int) *Atom {
  return &Atom{numberValue,val}
}

func NewSymbol(val string) *Atom {
  return &Atom{symbolValue,val}
}
