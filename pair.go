package littlelisp

type Pair struct {
  car evaluable
  cdr evaluable
}

func (p *Pair) eval(env *Env) string {
  if isNull(p.cdr) {
    return p.car.eval(env)
  } else if isNull(p.car) && isNull(p.cdr) {
    return ""
  } else {
    return p.car.evalList(env) + " " + p.cdr.eval(env)
  }
}

func (p *Pair) evalList(env *Env) string {
  if isSymbol(p.car) {
    return p.car.(*Symbol).Call(env, p.cdr)
  }

  return "(" + p.eval(env) + ")"
}

func cons(car evaluable, cdr evaluable) evaluable {
  return &Pair{car, cdr}
}

func emptyList() evaluable {
  return &Pair{null(), null()}
}
