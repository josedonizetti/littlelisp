package context

type Pair struct {
  car Value
  cdr Value
}

func (p *Pair) Eval(env *Env, forms *Forms) (Value, error) {
  if p.car == nil && p.cdr == nil {
    return nil, nil
  }

  if (IsSymbol(car(p))) {
    name := symbol(car(p))

    form := forms.Lookup(name)
    if form != nil {
      return form(p.cdr, env, forms), nil
    }

    procedure := env.Lookup(name)

    if procedure != nil {
      return procedure(p.cdr), nil
    }
  }

  value, _ := p.car.Eval(env, forms)
  return value, nil
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
