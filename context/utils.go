package context

func number(value Value) int {
  return value.(*Atom).Number()
}

func pair(value Value) *Pair {
  return value.(*Pair)
}

func car(value Value) Value {
  return pair(value).car
}

func cdr(value Value) *Pair {
  return pair(pair(value).cdr)
}

func NewAtom(val interface{},typ valueType) Value {
  return &Atom{typ,val}
}

func NewString(val string) Value {
  return &Atom{stringValue,val}
}

func NewNumber(val int) Value {
  return &Atom{numberValue,val}
}

func NewSymbol(val string) Value {
  return &Atom{symbolValue,val}
}
