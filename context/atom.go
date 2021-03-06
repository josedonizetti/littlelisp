package context

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

func (a *Atom) Eval(env *Env, forms *Forms) (Value, error) {
  switch a.typ {
  case symbolValue:
    return env.Lookup(a.val.(string))(nil), nil
  default:
    return a, nil
  }
}

func (a *Atom) String() string {
  switch a.typ {
  case symbolValue:
    fallthrough
  case stringValue:
    return a.val.(string)
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
  switch a.typ {
  case symbolValue:
    return true
  default:
    return false
  }
}
