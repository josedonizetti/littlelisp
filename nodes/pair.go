package littlelisp

type Pair struct {
  car *Atom
  cdr *Pair
}

func (p *Pair) Eval(env *Env) *Atom {
  return nil
}

func (p *Pair) String() string {
  return "(" + convertToString(p) + ")"
}

func convertToString(p *Pair) string {
  if p.car == nil && p.cdr == nil {
    return ""
  }

  if p.cdr == nil {
    return p.car.String()
  }

  return p.car.String() + " " + convertToString(p.cdr)
}

func NewPair(car *Atom, cdr *Pair) *Pair {
  return &Pair{car,cdr}
}

func EmptyPair() *Pair {
  return &Pair{nil,nil}
}
