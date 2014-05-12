package littlelisp

type Pair struct {
  car Value
  cdr Value
}

func (p *Pair) Eval(env *Env) (Value, error) {
  return nil, nil
}

func (p *Pair) String() string {
  if p.car == nil && p.cdr == nil {
    return "()"
  }

  return "(" + convertToString(p) + ")"
}

func convertToString(p *Pair) string {
  if p.car == nil && p.cdr == nil {
    return "()"
  }

  if p.cdr == nil {
    return p.car.String()
  }

  return p.car.String() + " " + convertToString(p.cdr.(*Pair))
}

func NewPair(car Value, cdr Value) *Pair {
  return &Pair{car,cdr}
}

func EmptyPair() Value {
  return &Pair{nil,nil}
}
