package context

func Add(value Value) Value {
  car := car(value)

  if pair(value).cdr == nil {
    return car
  }

  val := number(car) + number(Add(cdr(value)))
  return NewNumber(val)
}

func Mul(value Value) Value {
  car := car(value)

  if pair(value).cdr == nil {
    return car
  }

  val := number(car) * number(Mul(cdr(value)))
  return NewNumber(val)
}

func Div(value Value) Value {
  car := car(value)

  if pair(value).cdr == nil {
    return car
  }

  val := number(car) / number(Div(cdr(value)))
  return NewNumber(val)
}

func Sub(value Value) Value {
  car := car(value)

  if pair(value).cdr == nil {
    return car
  }

  val := number(car) + number(Sub(cdr(value)))
  return NewNumber(val)
}
