package context

type Pair struct {
  car Value
  cdr Value
}

func (p *Pair) Eval(env *Env) (Value, error) {
  name := symbol(car(p))

  procedure := env.Lookup(name)

  if procedure != nil {
    return procedure(p.cdr), nil
  }

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

  return p.car.String() + " " + convertToString(pair(p.cdr))
}

func NewPair(car Value, cdr Value) *Pair {
  return &Pair{car,cdr}
}

func EmptyPair() Value {
  return &Pair{nil,nil}
}
