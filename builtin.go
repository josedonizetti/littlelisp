package littlelisp

func Add(pair *Pair) *Atom {
  if pair.cdr == nil {
    return pair.car
  }
  val := pair.car.Number() + Add(pair.cdr).Number()
  return NewNumber(val)
}

func Mul(pair *Pair) *Atom {
  if pair.cdr == nil {
    return pair.car
  }
  val := pair.car.Number() * Add(pair.cdr).Number()
  return NewNumber(val)
}

func Div(pair *Pair) *Atom {
  if pair.cdr == nil {
    return pair.car
  }
  val := pair.car.Number() / Add(pair.cdr).Number()
  return NewNumber(val)
}

func Sub(pair *Pair) *Atom {
  if pair.cdr == nil {
    return pair.car
  }
  val := pair.car.Number() - Add(pair.cdr).Number()
  return NewNumber(val)
}
