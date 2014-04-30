package littlelisp

type Pair struct {
  car evaluable
  cdr evaluable
}

func (p *Pair) eval() string {
  if isNull(p.cdr) {
    return p.car.eval()
  } else if isNull(p.car) && isNull(p.cdr) {
    return ""
  } else {
    return p.car.evalList() + " " + p.cdr.eval()
  }
}

func (p *Pair) evalList() string {
  return "(" + p.eval() + ")"
}

func cons(car evaluable, cdr evaluable) evaluable {
  return &Pair{car, cdr}
}

func emptyList() evaluable {
  return &Pair{null(), null()}
}
